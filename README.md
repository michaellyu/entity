Entity
========

Shortcut to access interfaces.

### Installing
```shell
go get -u github.com/michaellyu/entity
```

### How To Use

```go
package main

import (
  "github.com/michaellyu/entity"
  "fmt"
)

func main() {
  var data interface{} = map[string]interface{} {
    "name": "Team",
    "city": map[string]interface{} {
      "name": "Tianjin",
    },
    "members": []interface{} {
      map[string]interface{} {
        "name": "Coral",
      },
      map[string]interface{} {
        "name": "Jack",
      },
      map[string]interface{} {
        "name": "Xinxing",
      },
    },
  }
  // get a interface{}
  name, err := team.Get("name")
  fmt.Println(name, err)
  // get a string
  cityName, _ := team.GetString("city.name")
  fmt.Println(cityName)
  // get a map
  firstMember, _ := team.GetMap("members[0]")
  fmt.Println(firstMember)
  // get a array
  members, _ := team.GetArray("members")
  fmt.Println(members)
  // set a field
  err = team.Set("name", "One Team")
  fmt.Println(err)
}
```

### Start Docker

```shell
sudo docker-compose up -d
```

### Run Test

```shell
./test
```

#### Run Benchmark

```shell
./test -bench="."
```

### End

```shell
sudo docker-compose down
```
