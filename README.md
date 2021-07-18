# message broker using etcd and go

- start etcd - `docker run --rm --name etcd -p 2379:2379 gcr.io/etcd-development/etcd /usr/local/bin/etcd --listen-client-urls http://0.0.0.0:2379 --advertise-client-urls http://192.168.99.100:2379`
- Change broker type between in-memory and etcd using `brokerType` flag
- Acceptable values for `brokerType` : `inmemory`(default), `etcd`