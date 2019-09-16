1. one channel per user to handle server updates, so that the docUpdater routine is not blocked on sending delta over websocket
2. look into http2/grpc/protobuf