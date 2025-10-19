
- server: c7i.xlarge
    - gp3 25GB 8000 IOPS
    - Ubuntu 24.04
    - mostly default postgres settings
        - log tables are configured to vacuum analyze aggressively
    - goal: ~60% CPU
    - us-east-1a

# REPLICATION
