# homelab
Personal homelab (services) running with Traefik and Docker

Here is the following structure of the project : 

```
|-- LICENSE
|-- README.md
|-- data
|   |-- fail2ban
|   |-- fbdata
|   |-- focalboard
|   |-- mealie-data
|   |-- pgdata
|   `-- vw-data
|-- docker-compose.yml
|-- letsencrypt
|   `-- acme.json
`-- logs
    `-- var
```

TL;DR : `docker-compose.yml` contains traefik configuration and `data/` contains all data for each services.
