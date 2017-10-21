# echo-sandbox

A sandbox for the [echo] framework.

## Usage

### With Go

```
$ make deps
$ make build
$ PORT=3000 build/echo-sandbox &
$ curl localhost:3000
```

### With Docker

```
$ docker build -t tsub/echo-sandbox .
$ docker run -d -p 3000:3000 tsub/echo-sandbox
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
