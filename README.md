# google-docs-clone

### Todo
1. define payload format for client to server and server to client communication
2. implement doc datastructure on top of pages
3. implement operational transformation for collaborative editing

#### Todo - immediate
1. channels and synchronization (mutex) revision
2. separate goroutine for updating document, listen on a buffered channel
3. user specific goroutines should send delta on the channel
4. make this work without parallel updates(without maintaining individual serverviewa) first, then move on to handle simultaneous edits