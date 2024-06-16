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
  $ docker run -p 3333:3333 go_lang_base_app:latest
`
