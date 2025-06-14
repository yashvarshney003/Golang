package main 
import (
	"fmt"
	"net/http"
	"bytes"
	"encoding/json"
	"strconv"
	"io/ioutil"
)
func main(){
	url := "http://localhost:8080/products/"
	for i:=0; i<10;i++ {

		data := map[string]interface{} {
			"Name" : "Products"  + strconv.Itoa(i),
			"Description": "This is practical second",
			"Quantity": 5,
			"Price": 4,
		}
		json_data, err := json.Marshal(data)
		if err!=nil{
			fmt.Println("Error in marshalling")
		}
		resp,err := http.Post(url,"application/json", bytes.NewBuffer(json_data))
		if err != nil{
			fmt.Println("error in sending post request", err)
		}
		defer resp.Body.Close()
		body,err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("❌ Error reading response body:", err)
			continue
		}
		fmt.Printf("✅ Product %d added, Status: %s, Response: %s\n", i, resp.Status, string(body))
		

	}
}