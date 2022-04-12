# regexp-rename
Rename files in a directory using regexp


## Build

```
go build
```

## Usage

To rename ./test/test 01.mp4 to ./test/hello 01.mp4

```
regexp-rename -dir ./test \
    -expr "(?m)^.*\s(\d*).(mp4)$" \
    -name-template "hello \${1}.\${2}" \
    -dry-run
```

Note: need to escape `$` for `name-template`