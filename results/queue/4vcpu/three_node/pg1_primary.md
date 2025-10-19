# Primary (pg1)

Here are the commands used to set up the primary:


1. Regular commands as detailed in here [README](../../../SETUP.md#postgres-setup)
2. Replication-specific commands
```bash

NODE_NAME=$(hostname -s)
REPL_USER=replicator
REPL_PASS=strongpass
```

```bash

   
sudo -u postgres psql -c "CREATE ROLE $REPL_USER WITH REPLICATION LOGIN PASSWORD '$REPL_PASS';"
# some configs are dubious but whatever
sudo tee -a /etc/postgresql/17/main/postgresql.conf > /dev/null <<'EOF'
fsync = on
synchronous_commit = on          # wait for WAL flush on commit
full_page_writes = on            # protect torn pages
wal_sync_method = open_datasync
commit_delay = 0                 # no batching delay
commit_siblings = 5              # default
wal_level = replica
max_wal_senders = 10
max_replication_slots = 10
wal_log_hints = on
archive_mode = off
hot_standby = on
synchronous_standby_names = 'FIRST 1 (pg2, pg3)'
wal_compression = on
hot_standby_feedback = on
max_worker_processes = 32
max_parallel_workers = 32
max_parallel_workers_per_gather = 16
EOF
   
sudo systemctl restart postgresql@17-main
# check if replication slots are taken
sudo -u postgres psql -c "SELECT * FROM pg_create_physical_replication_slot('pg2');"
sudo -u postgres psql -c "SELECT * FROM pg_create_physical_replication_slot('pg3');"

sudo -u postgres psql -c "SELECT application_name, client_addr, state, sync_state FROM pg_stat_replication;"
# explicitly enable replicator
echo "replication replicator 172.31.0.0/16 scram-sha-256" | sudo tee -a /etc/postgresql/17/main/pg_hba.conf
  
```

# Check Latest Replicated Status
```bash
sudo -u postgres psql -x -c "SELECT application_name, client_addr, state, sync_state, sent_lsn, flush_lsn, replay_lsn FROM pg_stat_replication;"
```