# Build instructions

```
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./public/
cd ./cmd/web/
GOOS=js GOARCH=wasm go build -o main.wasm . ; cp main.wasm ../../public
cd ../server
go run main.go
```


# TinyGo

Unfortunately tinygo can't compile most of the net/http packages

## WASM Exec JS

```
cp $(tinygo env TINYGOROOT)/targets/wasm_exec.js .
```


# Todo List Application with `elem-go`, `htmx`, and `Go labstack/echo`

Based off this [example](https://github.com/chasefleming/elem-go/tree/main/examples/htmx-fiber-todo) but with modifications. I grabbed this example as it looked like it did most of the todo stuff I wanted.

I did not realize what elem-go was - and so in the future I'll probably swap that out for a template based generation but most of the effort of this was towards doing the local first aspect.

