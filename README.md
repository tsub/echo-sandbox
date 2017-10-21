# echo-sandbox

A sandbox for the [echo] framework.

## Usage

### With Go

```
$ make dotenv
$ export $(cat .env)
$ make deps
$ make build
$ docker-compose up -d
$ build/echo-sandbox &
$ curl localhost:3000
```

## Deployments

### Heroku Container Registry

```
$ heroku container:login
$ heroku create
$ heroku container:push -a <app name> web
$ heroku open -a <app name>
```

[echo]: https://github.com/labstack/echo
