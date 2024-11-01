# Distributed Key-Value Store

This distributed key-value store is inspired by the [Google File System](https://static.googleusercontent.com/media/research.google.com/en//archive/gfs-sosp2003.pdf) and consists of three main components: Client Server, Master Server (Coordination Server), and Storage Node. Each component plays a crucial role in ensuring efficient storage and retrieval of key-value pairs across a distributed architecture.

## Client Server
The Client Server acts as the interface for clients to interact with the key-value store.

- ✅ Accepts and processes client requests via RESTful API.
- ✅ Communicates with other servers using gRPC for efficient data transfer.
- ✅ Routes storage and retrieval requests to the appropriate Storage Node based on guidance from the Master Server.

## Master Server (Coordination Server)
The Master Server coordinates interactions between clients and storage nodes, managing the overall architecture.

- ✅ Handles requests from Client Servers to determine the location of storage nodes.
- ✅ Responds to Client Servers with the IP and port of the designated Storage Node.
- ✅ Maintains metadata and state information about Storage Nodes for efficient load balancing.
- ✅ Facilitates the storage and retrieval workflow, ensuring consistency and availability.

## Storage Node
Storage Nodes are responsible for the actual storage of key-value pairs and maintain communication with the Master Server.

- ✅ Accepts incoming PUT and GET requests from Client Servers.
- ✅ Sends regular heartbeats to the Master Server to indicate operational status.
- ✅ Supports high availability and redundancy for stored data.
- ✅ Provides quick access to stored data through optimized retrieval mechanisms.

This key-value store architecture ensures scalability, reliability, and efficient data management across distributed systems.
