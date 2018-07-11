# Rest API server
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

`ypack` Rest API server.

## Requisites
To start hacking you need:
* [Golang 1.10](https://golang.org/)
* [Dep](https://github.com/golang/dep)
* [MariaDB](https://mariadb.org/)

## Build
Clone this repo and build it:

```bash
$ go get -v github.com/ypack/ypack-server
$ cd cd ypack-server
$ go build
```

Now you need to create the `ypack` database. When done, execute the [database.sql](database.sql)
script to create the database tables, indexes etc.

Before execute, you can change some parameters:

| Field     	| Description            	| Default value 	|
|-----------	|------------------------	|---------------	|
| `db-host` 	| Database host          	| `localhost`   	|
| `db-port` 	| Database port          	| `3306`        	|
| `db-name` 	| Database name          	| `ypack`       	|
| `db-user` 	| Database user          	|               	|
| `db-pass` 	| Database user password 	|               	|
| `host`    	| Server host            	| `localhost`   	|
| `port`    	| Server listen port     	| `8080`        	|
| `debug`    	| Show debug logs when needed     	| `false`        	|


## Start
You can start the server with all default values and provide the database user and password:

```bash
$ ypack-server -db-user=myuser -db-pass=mypassword
```
>**Important**: you need to have a running MariaDB instance with a database called `ypack`

If an error occurs, it will be shown.

## Test
You can execute all test using:

```bash
$ go test ../...
```
