package main 
import (
	"fmt"	
	"encoding/json"
)

func main(){
m1 := map[string]string{"key1": "value1","key2": "value2","key3": "value3"}

b, err := json.Marshal(m1)

fmt.Println(string(b))
if err != nil {
	fmt.Println("error:", err)
}
d1 :=map[string]string{"key": "value12","key2w": "value2q","kewy3": "value3"}
json.Unmarshal(b, d1)
fmt.Println(d1)
}
