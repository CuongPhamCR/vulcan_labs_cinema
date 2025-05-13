"# vulcan_labs_cinema" 

1. Prepare ENV
```bash
cd configs
touch config.yaml
cp local.yaml config.yaml
```

2. Start Local
```bash
go run cmd/main.go
```

3. Build & start by docker compose
```bash
docker compose up -d --build
```