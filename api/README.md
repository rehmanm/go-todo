# go-todo

```
go mod init rehmanm.go-todo
mkdir -p bin cmd/api internal migrations remote
touch Makefile
touch cmd/api/main.go
```

Health Check

```
curl http://localhost:4000/v1/healthcheck
```

## Migrations

## Create

```
migrate create -seq -ext=.sql -dir=./migrations {migration name}
```

## Execute

```
migrate -path=./migrations -database=$TODO_DB_SSN up
```

## Rollback

```
migrate -path=./migrations -database=$TODO_DB_SSN down
```
