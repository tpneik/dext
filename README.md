# dext

## setup 
```bash
go get -u github.com/spf13/cobra@latest

# then
go install github.com/spf13/cobra-cli@latest
```

```bash
cobra-cli init

cobra-cli add daemon
cobra-cli add daemon client
```
# Learning golang sytax with strings

```golang
// remember to import strings
import "strings"

// strings.Split("Your string", {your separator})
lines := strings.Split("My full string", " ")
// return a slice which is a dynamic array
// can be integrated with range , can be used like this

for _, line := range lines {}
// remember, range return index and value, if not use the index, 
// remember to assign it as blank identifier `_`, because golang not allow us 
// to declare a variable and not use

// strings.SplitN({your string | string variable}, {separator}, {number of part})
line := "Docker Root Dir: /var/lib/docker"
parts := strings.SplitN(line, "Docker Root Dir: ", 2 )
// part[0] --> "Docker Root Dir: "
// part[1] --> " /var/lib/docker"

// use strings.TrimSpace(part[0]) to clean the pre space
// strings.TrimSpace(part[1]) --> "/var/lib/docker"
```