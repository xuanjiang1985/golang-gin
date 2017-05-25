# golang-gin
#### A website application with golang gin framework

### clone this repository
```
cd $GOPATH/src
git clone https://github.com/xuanjiang1985/golang-gin
```
### install js deps from package.json
```
cd $GOPATH/src/golang-gin
npm install
```

### install golang deps from ./Godeps/Godeps.json. install godep firstly, if you have no godeps.

```
cd $GOPATH/src/golang-gin
$GOBIN/godep restore
```
### copy .env.example to .env and set mysql host.
`cp .env.example .env`

### run migrations to create database
```
cd $GOPATH/src/golang-gin/migrations
go run migrate.go
```
### finally
`go run main.go
`
#### localhost:8080
