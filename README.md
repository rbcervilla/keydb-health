# keydb-health
HTTP health check. KeyDB instance is ready to server requests if:
- Is not loading from persistence
- Is not synchronizing with master.

### Usage

```
Usage:
  -h string
    	KeyDB host (default "127.0.0.1")
  -p string
    	KeyDB port (default "6379")
  -sp string
    	Server listener port (default "8080")
```

### Docker

You can run it with Docker

```sh
docker run --name keydb-health rbcervilla/keydb-health -h keydbhost -p 6379
```