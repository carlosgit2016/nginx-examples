### Nginx load balancer examples

#### Simple round robin test

```bash
./build.sh --app serverapp # Build simple server app
./build.sh --app nginx # Build nginx image with custom configs
```

```bash
docker compose up --build # Up containers
```

Execute this command more than once to get different container ip and HOSTNAME
```bash
# Testing balancer
curl -H "Host: serverapp.com" localhost:3000
{
    "ip": "192.168.16.3", # The ip of the container extracted from /etc/hosts
    "container_name": "504d4643d32f" # The HOSTNAME of the container
}
```