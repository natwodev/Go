# PostgreSQL Setup and Connection

## Install PostgreSQL
- Use Docker: `docker run --name go-postgres -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres`
- Or install locally on MacOS using Homebrew: `brew install postgresql`

## Connection Configuration
Configuring Golang to connect to PostgreSQL using `database/sql` and `github.com/lib/pq`.

### Sample DSN
`host=localhost port=5432 user=postgres password=mysecretpassword dbname=godb sslmode=disable`
