# ligo
Useful go library

### KeySearch
Note: All JSON numbers will be returned as float64
```go
byteData := "{ \"number\": 123, \"object\": { \"child\": \"hello world!\" } }"
var data interface{}
err := json.Unmarshal([]byte(byteData), &data)
integer := KeySearch("number", data, false).(float64)
fmt.Println(integer) // 123
```
