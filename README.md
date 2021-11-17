# bandwidth-trasher

Simple TCP application layer throughput tester.  It simulates an SSL session by encrypting an empty payload and sending this over the wire.

## Server

When ran in server mode, it will listen on a port and wait for incoming connections.
```bash
$ ./server
```

## Sender

When ran in sender mode, it will dial out and then push data to the server port.
```bash
$ ./sender
```

## Puller

When ran in puller mode, it will dial out and then listen for the server to back.
```bash
$ ./puller
```
