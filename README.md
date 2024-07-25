### Build docker image
`
  $ docker build -t go_lang_base_app:latest -f Dockerfile.dev .
`

### Start DB
`
  $ docker-compose up
`

### Start Docker container
`
  $  docker build -t  go_lang_base_app:latest  -f Dockerfile.dev .
`


`
  $ docker run -p 3333:3333 go_lang_base_app:latest
`

### Using Docker Compose

`
  $ docker compose up
`

* If you use only DB (On windows, facing an issue while using docker compose with air, so have to start dev server separately)
`
  $ docker compose up postgres
`

* Start app server individually in the same network

`
  $ docker run -p 3333:3333 --network go-lang-base-app_default go_lang_base_app:latest
`

### Notes
* When connecting from app server to db server, the connection url should point to service host, not localhost
`postgresql://root:secret@postgres:5432/go_app_database?sslmode=disable`

Here postgres is name of the service in docker-compose.yml

* air not working as expected on windows

#### Windows solution
* Don't run web server in docker mode, run directly
* generate windows .air.toml separately on windows
* `$ air init`
* `$ air http_server`
*  Connection string `postgresql://root:secret@localhost:5433/go_app_database?sslmode=disable`