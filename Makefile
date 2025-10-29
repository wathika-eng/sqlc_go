migrate:
	@dbmate up

run:
	@go run .

build:
	@go build -o sqlc_go .

clean:
	@rm -f sqlc_go

gen:
	@sqlc generate