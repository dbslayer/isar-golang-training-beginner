1. Prepare .env from .env.example
   ```
   cp .env.example .env
   ```
   update file .env accordingly
2. Create docker from Dockerfile and
   ```
   docker build -t postgres-db:latest -f Dockerfile .
   ```
3. run docker
   ```
   docker compose up
   ```

4. run application
   ```
   go run server.go
   ```