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

## Work with git:

| Prefix | Meaning | Example |
|--------|---------|---------|
| feature | Some new functions | feature: New algorithm |
| fix | Fixing some bugs/errors/panics | fix: Buffer overflow |
| add | Added new file or directory | add: go.mod |
| refactor | Renamed/moved/removed some files or dirs | refactor: Renamed abc.go to bcd.go |