- Operates three unique services, in three different runtimes, each exposing
  both a unary RPC (from previous example) and a full-duplex streaming RPC
- Services are segragated into three isolated networks with at most two services
  per network
- Services are daisy-chained to each other via their streaming RPCs
- Python client accesses single service on first network and gets a reply from
  service in furthermost network
