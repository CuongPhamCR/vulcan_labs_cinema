"# vulcan_labs_cinema" 

1. Prepare ENV
```bash
cd configs
touch config.yaml
cp local.yaml config.yaml
```

2. Install package
```bash
go mod tidy
```


3. Start Local
```bash
go run cmd/main.go
```

4. Build & start by docker compose
```bash
docker compose up -d --build
```