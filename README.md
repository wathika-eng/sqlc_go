
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


```
