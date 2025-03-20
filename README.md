# homelab
Personal homelab (services) running with Traefik and Docker

Here is the following structure of the project : 

```
|-- LICENSE
|-- README.md
|-- data
|   |-- service1-data
|   |-- service2-data
|   |-- service3-data
|   |-- (...)
|-- docker-compose.yml
|-- letsencrypt
|   `-- acme.json
`-- logs
    `-- service1-logs
    `-- service2-logs
    `-- service3-logs
    `-- (...)
```

TL;DR : `docker-compose.yml` contains traefik configuration, `data/` and `logs` contains all data / logs for each services.
