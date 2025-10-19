⚠️ NOTE: I fiddled around with these settings a lot. The exact sequence of commands may be slightly different than the one I ran. The main issues I encountered was from connecting replicas to each other.

1. Regular commands as detailed in here [README](../../../SETUP.md#postgres-setup)
2. Replication-specific commands for pg3 as the follower are:
```bash
PG1_IP=172.31.95.101
REPL_USER=replicator

# start replication as pg2
sudo -u postgres pg_basebackup -h $PG1_IP -D $MNT/pg17 -U $REPL_USER -P -R --slot=pg2

sudo cp /usr/share/postgresql/17/postgresql.conf.sample /pgdata/pg17/postgresql.conf
sudo cp /usr/share/postgresql/17/pg_hba.conf.sample /pgdata/pg17/pg_hba.conf
sudo cp /usr/share/postgresql/17/pg_ident.conf.sample /pgdata/pg17/pg_ident.conf

sudo tee -a /pgdata/pg17/postgresql.conf > /dev/null <<'EOF'
port = 5432
listen_addresses = '*'
hot_standby = on
# match primary settings to allow recovery
max_worker_processes = 32
max_parallel_workers = 32
max_parallel_workers_per_gather = 16
max_wal_senders = 10
max_replication_slots = 10
EOF


sudo tee /pgdata/pg17/pg_hba.conf > /dev/null <<'EOF'
# TYPE  DATABASE        USER            ADDRESS                 METHOD

# allow local connections for postgres user
local   all             postgres                                peer

# allow replication connections from the primary
host    replication     replicator      172.31.0.0/16           scram-sha-256

# allow local TCP
host    all             all             127.0.0.1/32            scram-sha-256
host    all             all             ::1/128                 scram-sha-256
EOF

sudo chown postgres:postgres /pgdata/pg17/*.conf
sudo -u postgres /usr/lib/postgresql/17/bin/postgres -D /pgdata/pg17 -c config_file=/pgdata/pg17/postgresql.conf
sudo grep application_name /pgdata/pg17/postgresql.auto.conf || echo "primary_conninfo = 'user=replicator password=strongpass host=172.31.95.101 port=5432 application_name=pg3'" | sudo tee /pgdata/pg17/postgresql.auto.conf
sudo -u postgres /usr/lib/postgresql/17/bin/postgres -D /pgdata/pg17 -c config_file=/pgdata/pg17/postgresql.conf
```

The end result should look something like this:
```bash
sudo -u postgres psql -x -c "SELECT pg_is_in_recovery();"

Ver Cluster Port Status Owner Data directory Log file
-[ RECORD 1 ]-----+--
pg_is_in_recovery | t

# confirm connection string and slot name
sudo -u postgres grep -E 'primary_conninfo|primary_slot_name' /pgdata/pg17/postgresql.auto.conf

primary_conninfo = 'user=replicator password=strongpass host=172.31.95.101 port=5432 application_name=pg2'
primary_slot_name = 'pg2'

# confirm streaming status from standby perspective
sudo -u postgres psql -x -c "SELECT status, sender_host, sender_port, conninfo FROM pg_stat_wal_receiver;"

-[ RECORD 1 ]--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
status      | streaming
sender_host | 172.31.95.101
sender_port | 5432
conninfo    | user=replicator password=******** channel_binding=prefer dbname=replication host=172.31.95.101 port=5432 application_name=pg2 fallback_application_name=walreceiver sslmode=prefer sslnegotiation=postgres sslcompression=0 sslcertmode=allow sslsni=1 ssl_min_protocol_version=TLSv1.2 gssencmode=prefer krbsrvname=postgres gssdelegation=0 target_session_attrs=any load_balance_hosts=disable

# check local LSN vs replayed
sudo -u postgres psql -x -c "SELECT pg_last_wal_receive_lsn(), pg_last_wal_replay_lsn(), pg_is_wal_replay_paused();"

-[ RECORD 1 ]-----------+-----------
pg_last_wal_receive_lsn | 0/E1C4FEB8
pg_last_wal_replay_lsn  | 0/E1C4FEB8
pg_is_wal_replay_paused | f 
```
