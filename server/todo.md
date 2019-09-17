1. one channel per user to handle server updates, so that the docUpdater routine is not blocked on sending delta over websocket
2. look into http2/grpc/protobuf
3. two seperate bitmaps? one for insertion, one for deletion to save memory.

## peristence
1. ask for username in ui
2. use that as user id
3. disable editor while offline, reconnect capability
4. define startup sequence,i.e set of messages exchanged between server and client to sync up. 
5. instead of removing a user mark that user as inactive, use in memory hashmap for username to user struct mapping, useful when user comes back online after inactivity
6. scaling to multiple docs -- doc id, in memmory hashmap from doc to doc structure


## ui
1. go offline button in ui
2. sequence numbers, send last incorporated delta's sequence number along with delta updates


# package
1. create new package doc, move page package inside doc package