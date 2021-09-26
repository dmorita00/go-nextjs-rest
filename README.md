# docker-compose

before docker up, add database env

```shell
$ touch docker/db/env
$ vi docker/db/env
TZ=Asia/Tokyo
MYSQL_DATABASE=example
MYSQL_USER=example
MYSQL_PASSWORD=example
MYSQL_ALLOW_EMPTY_PASSWORD=yes
MYSQL_ROOT_PASSWORD=example
```

```shell
$ docker-compose up -d --build
```

# Go

```shell
$ docker-compose exec server bash
$ go install
```

## Run server

- normal mode
```shell
$ go run main.go
```
- debug mode
```shell
$ air -c .air.toml
```

### If in debug mode, watch the port

#### By JetBrains

- Add Go remote
  - host: localhost
  - port: 12345

# Next.js

```shell
$ docker-compose exec client bash
$ yarn
$ yarn dev
```

# MySQL

host:     localhost
port:     13306
database: example
user:     example
password: example
