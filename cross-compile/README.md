# Golang Cross-Compile

> Cross Compile Command

```
   go tool dist list

   // Linux -> Windows
   // Mac -> Windows
   GOOS=windows GOARCH=amd64 go build -o myprogram.exe main.go

   // Windows -> Linux
   SET GOOS=linux
   SET GOARCH=amd64
   go build -o myprogram main.go
```

> use Docker Command

```
    // deploy use multi staging step (10M ~ 20M)
    docker build -f Dockerfile -t golang .


    // deploy use simple step (300M ~ 400M)
    docker build -f Dockerfile.notStage -t golang .
```
