# gator

`gator` is a CLI RSS aggregator written in Go.

It lets you:
- register/login users
- add and follow feeds
- scrape RSS feeds on an interval
- browse recent posts from feeds you follow

## Requirements

You need these installed locally:
- **Go** (for `go install` / building)
- **PostgreSQL** (for storing users, feeds, follows, and posts)
- **Goose** (for database migrations)

## Install the CLI

Install directly from GitHub:

```bash
go install github.com/DennisPrudlik/gator@latest
```

After install, make sure your Go bin directory is on `PATH` (commonly `$HOME/go/bin`).

## Configure gator

Create a config file at:

```bash
~/.gatorconfig.json
```

Example:

```json
{
  "db_url": "postgres://YOUR_USER:YOUR_PASSWORD@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}
```

## Setup database

From the project root, run migrations:

```bash
goose -dir sql/schema postgres "postgres://YOUR_USER:YOUR_PASSWORD@localhost:5432/gator?sslmode=disable" up
```

## Run commands

For development, you can run with:

```bash
go run . <command> [args...]
```

For production use, run the compiled CLI binary (`gator`) installed via `go install`.

### Common commands

```bash
gator register <username>
gator login <username>

gator addfeed <feed_name> <feed_url>
gator follow <feed_url>
gator unfollow <feed_url>
gator following

gator feeds
gator users

gator agg 1m
gator browse
gator browse 10
```

## Notes

- `go run .` is intended for development.
- `gator` (the installed/built binary) is intended for normal usage.
- Go binaries are statically compiled executables, so once built/installed you can run `gator` directly.
