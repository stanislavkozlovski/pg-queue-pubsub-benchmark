To change the configs, I ran the following commands. They consist of:
1. Commenting out the settings we're overriding in the main `postgresql.conf` file
2. Adding the override settings in a `conf.d` override file.

# 1. Comment out default postgres settings

Create a bash file:
```bash
vi comment_out.sh
```
```bash
#!/bin/bash
CONF="/etc/postgresql/17/main/postgresql.conf"

# list of params to comment out
PARAMS=(
  max_connections
  superuser_reserved_connections
  shared_buffers
  effective_cache_size
  work_mem
  maintenance_work_mem
  huge_pages
  wal_level
  wal_compression
  wal_buffers
  min_wal_size
  max_wal_size
  checkpoint_timeout
  checkpoint_completion_target
  synchronous_commit
  default_statistics_target
  random_page_cost
  seq_page_cost
  effective_io_concurrency
  max_worker_processes
  max_parallel_workers
  max_parallel_workers_per_gather
  max_parallel_maintenance_workers
  autovacuum
  autovacuum_max_workers
  autovacuum_naptime
  autovacuum_vacuum_cost_limit
  autovacuum_vacuum_cost_delay
  jit
  track_io_timing
  log_checkpoints
  log_autovacuum_min_duration
  log_min_duration_statement
  shared_preload_libraries
  pg_stat_statements.max
  pg_stat_statements.track
)

for p in "${PARAMS[@]}"; do
  sudo sed -i "s/^\s*\(${p}\s*=\s*.*\)$/# \1/" "$CONF"
done

echo "✅ All matching parameters commented out in $CONF"

```
Run it.
```bash
chmod +x comment_out.sh
./comment_out.sh 
```

# 2. Paste the [modified PostgreSQL config](./modified_postgresql.conf) into a file in `conf.d`:
```bash
vi /etc/postgresql/17/main/conf.d/99-custom.conf
# paste
```

# 3. Restart PG
```bash
sudo systemctl reload postgresql@17-main
```

# 4. Check if settings are present
```bash
# psql into machine
SHOW shared_buffers;
```

# 5. Disable Transparent Huge Pages in the Kernel
```bash
echo never | sudo tee /sys/kernel/mm/transparent_hugepage/enabled
echo never | sudo tee /sys/kernel/mm/transparent_hugepage/defrag

sudo tee /etc/systemd/system/disable-thp.service >/dev/null <<'EOF'
    [Unit]
    Description=Disable Transparent Huge Pages
    DefaultDependencies=no
    Before=sysinit.target
    [Service]
    Type=oneshot
    ExecStart=/usr/bin/bash -lc 'echo never > /sys/kernel/mm/transparent_hugepage/enabled'
    ExecStart=/usr/bin/bash -lc 'echo never > /sys/kernel/mm/transparent_hugepage/defrag'
    RemainAfterExit=yes
    [Install]
    WantedBy=sysinit.target
EOF
```

# 6. Modify Kernel hugepages & allow Postgres to use them
```bash
sudo sysctl -w vm.nr_hugepages=24576
echo 'vm.nr_hugepages=24576' | sudo tee /etc/sysctl.d/60-postgres-hugepages.conf
gid=$(getent group postgres | cut -d: -f3)
echo "vm.hugetlb_shm_group=$gid" | sudo tee /etc/sysctl.d/61-postgres-hugetlb-group.conf
echo 27000 | sudo tee /sys/kernel/mm/hugepages/hugepages-2048kB/nr_hugepages
echo "vm.nr_hugepages = 27000" | sudo tee -a /etc/sysctl.conf
sudo sysctl --system
grep -i huge /proc/meminfo # check
```

# 7. Run the test with the option to tune the table

```bash
./pg_msg_bench \ # # [...general test flags...]
  --tune-table-vacuum # <--NEW
```
