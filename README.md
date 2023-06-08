# show command build

`boilerplate`

> initiate go module

```powershell
go mod init github.com/eduardkh/show
```

> get the cobra module and the cobra CLI tool

```powershell
go get -u github.com/spf13/cobra@latest
go install github.com/spf13/cobra-cli@latest
```

> initiate cobra project (must be the same as the go module show)

```powershell
cobra-cli init
```

> test the app

```powershell
go run main.go
go buld .
```

`functionality`

> add commands

```powershell
cobra-cli add ip
cobra-cli add interface -p ipCmd
```

> add gateway command

```powershell
cobra-cli add gateway -p ipCmd
go mod tidy
build.bat
```

> add external command

```powershell
cobra-cli add external -p ipCmd
go mod tidy
build.bat
```

> add calc command

```powershell
cobra-cli add calc -p ipCmd
go mod tidy
build.bat
```

> add brief command

```powershell
cobra-cli add brief -p interfaceCmd
build.bat
```

> add timestamp

```powershell
cobra-cli add timestamp
build.bat
```

## show command usage

> install the tool from github.com

```powershell
go install github.com/eduardkh/show@latest
```

> install autocompletion

```powershell
show completion powershell | Out-String | Invoke-Expression
```

> make autocompletion permanent in PS

```powershell
# in $PROFILE file
Test-Path $PROFILE
New-Item -path $PROFILE -type file -force
echo "show completion powershell | Out-String | Invoke-Expression" >> $PROFILE

Get-ExecutionPolicy
Set-ExecutionPolicy RemoteSigned # as admin

# revert $PROFILE file and policy
Remove-Item $PROFILE
Set-ExecutionPolicy Restricted # as admin
```

> basic usage

```powershell
show ip external
# Your external IP is: [public ip]

show ip interface brief
# IP Address      Subnet Mask     MAC Address             IP Enabled      Interface Description
# 192.168.7.50    255.255.255.0   04:7C:16:00:00:00       true            Intel(R) Ethernet Controller (3) I225-V

```
