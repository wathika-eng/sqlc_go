-- migrate:up transaction:false
CREATE DATABASE sqlc_go;

-- migrate:down
DROP DATABASE IF EXISTS sqlc_go;