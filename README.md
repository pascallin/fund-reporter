# fund reporter

Trying to crawler data of Chinese fund to analyze and visulize.

## Development

### install packages

```shell
git config --local core.hooksPath .githooks/ # optional, set githooks path
go mod download
```

### Try running

```shell
go run ./reporter.go crawl stat gdp
go run ./reporter.go crawl stat cpi
go run ./reporter.go crawl stat ppi
go run ./reporter.go crawl fund [code]

go run ./reporter.go save fund 161725 481010

go run ./reporter.go crawl stat gdp -t
```
