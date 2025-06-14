package main 
import (
	"fmt"
	"net/http"
	"io/ioutil"
)
func main(){
	url := "http://localhost:8080/products/"
	for i:=0; i<200;i++ {
		resp,err := http.Get(url)
		if err != nil{
			fmt.Println("error in sending get request", err)
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