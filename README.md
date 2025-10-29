
Ensure Go 1.23+ and Postgresql is installed, else if not:

```bash
bash <(curl -sL https://git.io/go-installer)

go version # go 1.25

curl -sSfL https://raw.githubusercontent.com/air-verse/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin # for hot reloading web server (optional)
psql --version # psql (PostgreSQL) 17
```

Setting up locally:

```bash
git clone
cd sqlc_go

go mod tidy # to install dependecies

cp .env.example .env # to create a local copy of the .env

#install dbmate using npm
npm install -g dbmate

# install sqlc
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

sqlc compile && sqlc vet # to verify and vet sqlc files, should be no errors

dbmate up # to run migrations and create the database, reads from .env

go run . # to seed the database and start the server
```
