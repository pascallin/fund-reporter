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

go run ./reporter.go save fund [code1] [code2]

go run ./reporter.go crawl stat gdp -t
```

### Help

```shell
> go run ./reporter.go help
获取中国基金及经济数据的命令行工具。
致在收集数据，可视化，并以文件格式进行导出方便做其他数据分析。

Usage:
  fund-reporter [command]

Available Commands:
  crawl       抓取数据
  help        Help about any command
  save        基金数据另存为

Flags:
  -h, --help   help for fund-reporter

Use "fund-reporter [command] --help" for more information about a command.
```

### Deployment

```shell
go build -o ./bin
```
