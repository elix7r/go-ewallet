## How to run
Server:
```console
    $foo@bar:~$ docker-compose up
    $foo@bar:~$ go run cmd/server/main.go
```

Client:
- 1: Create a new transaction
```console
    $foo@bar:~$ go run cmd/client/main.go -send cac9a790-1d76-4b24-a508-d3194cb73d0c c18e5639-4849-41c4-b34c-9a01a75720ef 650
```

- 2: Get a not issued transactions

```console
    $foo@bar:~$ go run cmd/client/main.go -getlast
```

### Existing wallets credentials
- cac9a790-1d76-4b24-a508-d3194cb73d0c (balance: 11350)
- c18e5639-4849-41c4-b34c-9a01a75720ef (balance: 4150)