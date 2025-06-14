package main 

import (
	"fmt"
	"net/http"
	"encoding/json"
	"bytes"
	"io/ioutil"
	"strconv"
	
)

func main(){
	url:= "http://localhost:8080/products/"
	for i:=0;i<10;i++{
	data := map[string]interface{}{
            "Name":        "Products" + strconv.Itoa(i),
            "Description": "This is a products one",
            "Price":       300,
            "Quantity":    4,
        }
	jsonData, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Request Error:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp,err := client.Do(req)
	if err!=nil{
		fmt.Println("Http err", err)

	}
	defer resp.Body.Close()
	body,_ :=  ioutil.ReadAll(resp.Body)
	fmt.Println("Response:",string(body))
} 

}