Entity
========
[![Go Report Card][go-report-image]][go-report-url]
[![GoDoc][godoc-image]][godoc-url]

[go-report-image]: https://goreportcard.com/badge/github.com/michaellyu/entity
[go-report-url]: https://goreportcard.com/report/github.com/michaellyu/entity
[godoc-image]: https://godoc.org/github.com/michaellyu/entity/lint?status.svg
[godoc-url]: https://godoc.org/github.com/michaellyu/entity

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
  // get an interface{}
  name, err := team.Get("name")
  fmt.Println(name, err)
  // get a string
  cityName, _ := team.GetString("city.name")
  fmt.Println(cityName)
  // get a map
  firstMember, _ := team.GetMap("members[0]")
  fmt.Println(firstMember)
  // get an array
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
