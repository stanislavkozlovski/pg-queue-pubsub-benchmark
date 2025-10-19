# Setup

A few useful series of commands to set up Postgres and the test

## Postgres Setup

1. Create a Server EC2. Pick the appropriate instance carefully
   - Ensure you open said instance's postgresql port via a security group
   - Pick Ubuntu 24.04 as the OS
2. Create and attach an EBS drive to the instance
3. Mount the EBS and install Postgres
```bash
export PRIVATE_IP="xxx" # the private IP of your instance


### ------ Mount the EBS drive to be used as Postgres' data directory ------

lsblk # to find the drive; it's usually /dev/nvme1n1
DEV=/dev/nvme1n1
MNT=/pgdata

sudo mkfs.xfs -f $DEV
sudo mkdir -p $MNT
sudo mount $DEV $MNT
sudo mkdir -p $MNT/pg17
UUID=$(sudo blkid -s UUID -o value $DEV)
echo "UUID=$UUID $MNT xfs defaults,nofail 0 2" | sudo tee -a /etc/fstab

### ------ Install PostgreSQL 17 (PGDG repo) ------
sudo apt-get update -y
sudo apt-get install -y curl ca-certificates gnupg lsb-release
curl https://www.postgresql.org/media/keys/ACCC4CF8.asc | \
  gpg --dearmor | sudo tee /etc/apt/trusted.gpg.d/pgdg.gpg >/dev/null
echo "deb http://apt.postgresql.org/pub/repos/apt $(lsb_release -cs)-pgdg main" | \
  sudo tee /etc/apt/sources.list.d/pgdg.list
sudo apt-get update -y # again after updating the sources
sudo apt-get install -y postgresql-17 postgresql-client-17 postgresql-common

### ------ Create the PG instance on the mounted volume
sudo pg_dropcluster --stop 17 main
sudo chown -R postgres:postgres $MNT/pg17
sudo chmod 700 $MNT/pg17
sudo chown postgres:postgres $MNT
sudo chmod 700 $MNT
sudo pg_createcluster 17 main --datadir=$MNT/pg17 --port=5432 -- -c listen_addresses="${PRIVATE_IP},localhost"
sudo systemctl start postgresql@17-main

pg_lsclusters # check that the cluster is up & running

### Modify the db-level settings to allow all external traffic, create the right benchmark DB and user
echo "host all all 0.0.0.0/0 md5" | sudo tee -a /etc/postgresql/17/main/pg_hba.conf
sudo systemctl reload postgresql@17-main
sudo -u postgres createdb benchmark
sudo -u postgres psql -c "ALTER USER postgres WITH PASSWORD 'postgres';"
```

## Client Driver Setup

1. Create a Client EC2, in the same AZ as the server.
2. Install Go
```bash
cd /tmp
wget https://go.dev/dl/go1.24.7.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.24.7.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
go version
```

3. Build the Benchmark
```bash
git clone https://github.com/stanislavkozlovski/pg-queue-pubsub-benchmark.git
cd pg-queue-pubsub-benchmark

go build -o ./pg_msg_bench .
chmod +x ./pg_msg_bench
```

4. Run the Benchmark

```bash
export HOST="xxx"  # adjust to your server's private IP

./pg_msg_bench \
  --host=$HOST \
  --port=5432 \
  --db=benchmark \
  --user=postgres \
  --password=postgres
  # [...other flags as needed...]
# See each specific tests' page for the exact command that was ran.
```

See the [General README](../README.md#usage) for the full client usage and flags.
