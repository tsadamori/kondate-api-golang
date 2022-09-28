# kondate-api-golang

## create migration file

```bash
migrate create -ext sql -dir database/migrations -seq [file name]
```

## exec migrations

```bash
migrate -path database/migrations -database 'mysql://[user]:[password]@tcp([host]:[port])/[db_name]' [up/down]
```
