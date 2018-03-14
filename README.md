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
    "author": map[string]interface{} {
      "id": 1,
      "name": "Coral",
      "children": [] int { 1, 2, 3 },
    },
  }
  entity := entity.NewEntity(&data)
  authorName, err := entity.GetString("author.name")
  if err != nil {
    fmt.Println("author.name not found")
  }
  fmt.Println("author.name:", authorName)
  firstChild, err := entity.GetInt("author.children[0]")
  if err != nil {
    fmt.Println("author.children[0] not found")
  }
  fmt.Println("author.children[0]:", firstChild)
  children, err := entity.GetArrayInt("author.children")
  if err != nil {
    fmt.Println("author.children not found")
  }
  fmt.Println("author.children:", children)
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
