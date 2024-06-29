# go-mservice

## Introducion:
This is a simple client-server application for my education.

## Architecture:
This is the target architecture of the project:


```
[client] <-- GRPC--> [server]
    ^                   |
    |                   |
    |                   ???
    by HTTP             |
    |                   |
    |                   psql DB
    Me                  
```