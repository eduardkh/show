# show

`boilerplate`

> initiate go module

```bash
go mod init github.com/eduardkh/show
```

> get the cobra module and the cobra CLI tool

```bash
go get -u github.com/spf13/cobra@latest
go install github.com/spf13/cobra-cli@latest
```

> initiate cobra project (must be the same as the go module show)

```bash
cobra-cli init
```

> test the app

```bash
go run main.go
go buld .
```

`functionality`

> add commands

```bash
cobra-cli add ip
cobra-cli add interface -p ipCmd
```

> add gateway command

```bash
cobra-cli add gateway -p ipCmd
go mod tidy
build.bat
```

> add brief command

```bash
cobra-cli add brief -p interfaceCmd
build.bat
```
