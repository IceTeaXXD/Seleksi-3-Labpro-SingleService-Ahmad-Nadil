## Migrate
```
go run migrate/migrate.go
```

## Compile
```
CompileDaemon -command="./singleservice"
```

## Build Docker
### 1. Build the docker file
```
docker build . -t singleservice:latest 
```

### 2. Run the docker (you can change the port accordingly)
```
docker run -e PORT=9000 -p 9000:9000 singleservice:latest
```