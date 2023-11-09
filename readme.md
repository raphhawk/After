## []<-[After]<-[]

### A Distributed Commit Log

The commit log takes in published message and stores in an Append only store
The published messages can be consumed and further used (as Events or States) in a stream.
The client-server communication is implemented with gRPC and helps with streaming of messages.
