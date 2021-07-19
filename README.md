# message broker using etcd and go

- start etcd - `brew install etcd`
- Change broker type between in-memory and etcd using `brokerType` flag
- Acceptable values for `brokerType` : `inmemory`, `etcd`(default)