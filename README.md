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

* If you use only DB
`
  $ docker compose up postgres
`

* Start app server individually in the same network

`
  $ docker run -p 3333:3333 --network go-lang-base-app_default go_lang_base_app:latest
`