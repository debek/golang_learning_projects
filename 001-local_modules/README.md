`go mod init example/username/repo` - module inicialize

```
❯ go test
PASS
ok  	github.com/debek/golang_learning_projects/001-local_modules	0.146s

❯ go mod init example.com/username/repo
go: creating new go.mod: module example.com/username/repo
go: to add module requirements and sums:
    go mod tidy

❯ go test
PASS
ok  	example.com/username/repo	0.132
```

Test pointer changed.
