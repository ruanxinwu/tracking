package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	//"net/url"
	"strings"
	//"encoding/json"
)

func main() {
	//1 List all carriers
	var url string = "http://api.trackingmore.com/v2/carriers/"
	var postData string = ""
	httpDo(url, postData, "GET")

	//2 detect a carriers by tracking number
	//var url string ="http://api.trackingmore.com/v2/carriers/detect"
	//var postData string ="{\"tracking_number\":\"EA152563251CN\"}"
	//httpDo(url,postData,"POST")

	//3 create a tracking number
	//var url string ="http://api.trackingmore.com/v2/trackings/post"
	//var postData string ="{\"tracking_number\": \"BYS006086078\",\"carrier_code\":\"yanwen\",\"title\":\"chase chen\",\"customer_name\":\"chase\",\"customer_email\":\"abc@qq.com\",\"order_id\":\"#123\",\"order_create_time\":\"2018-05-20 12:00\",\"destination_code\":\"IL\",\"tracking_ship_date\":\"1521314361\",\"tracking_postal_code\":\"13ES20\",\"lang\":\"en\",\"logistics_channel\":\"4PX page\"}"
	//httpDo(url,postData,"POST")

	//4 List all trackings
	//var url string ="http://api.trackingmore.com/v2/trackings/get?page=1&limit=100&created_at_min=1521314361&created_at_max=1541314361&update_time_min=1521314361&update_time_max=1541314361&order_created_time_min=&order_created_time_max=&numbers=&orders=&lang=en"
	//var postData string =""
	//httpDo(url,postData,"GET")

	//5 Get tracking results of a single tracking
	//var url string ="http://api.trackingmore.com/v2/trackings/yanwen/BYS006086078/cn"
	//var postData string =""
	//httpDo(url,postData,"GET")

	//6 create muti tracking number
	//var url string ="http://api.trackingmore.com/v2/trackings/batch"
	//var postData string = "[{\"tracking_number\": \"EA152565249CN\",\"carrier_code\":\"china-ems\",\"title\":\"chase chen\",\"customer_name\":\"chase\",\"customer_email\":\"abc@qq.com\",\"order_id\":\"#123444\",\"order_create_time\":\"2018-05-20 12:00\",\"destination_code\":\"IL\",\"tracking_ship_date\":\"1525314361\",\"tracking_postal_code\":\"13ES20\",\"lang\":\"en\",\"logistics_channel\":\"4PX page\"},{\"tracking_number\": \"EA152563246CN\",\"carrier_code\":\"china-ems\",\"title\":\"chase chen\",\"customer_name\":\"chase\",\"customer_email\":\"abc@qq.com\",\"order_id\":\"#123444\",\"order_create_time\":\"2018-05-20 12:00\",\"destination_code\":\"IL\",\"tracking_ship_date\":\"1521314361\",\"tracking_postal_code\":\"13ES20\",\"lang\":\"en\",\"logistics_channel\":\"4PX page1\"}]"
	//httpDo(url,postData,"POST")

	//7 Update Tracking item
	//var url string ="http://api.trackingmore.com/v2/trackings/yanwen/BYS006086078"
	//var postData string = "{\"title\": \"testtitle\",\"customer_name\":\"go test\",\"customer_email\":\"942632688@qq.com\",\"order_id\":\"#1234567\",\"logistics_channel\":\"chase chen go\",\"customer_phone\":\"+86 13873399982\",\"destination_code\":\"US\",\"status\":\"7\"}"
	//httpDo(url,postData,"PUT")

	//8 Delete a tracking item
	//var url string ="http://api.trackingmore.com/v2/trackings/yanwen/BYS006086078"
	//var postData string = ""
	//httpDo(url,postData,"DELETE")

	//9 Get realtime tracking results of a single tracking
	//var url string ="http://api.trackingmore.com/v2/trackings/realtime"
	//var postData string = "{\"tracking_number\": \"61290983300030854514\",\"carrier_code\":\"fedex\",\"destination_code\":\"US\",\"tracking_ship_date\": \"20180226\",\"tracking_postal_code\":\"13ES20\",\"specialNumberDestination\":\"US\",\"order\":\"#123123\",\"order_create_time\":\"2018/3/27 16:51\",\"lang\":\"en\"}"
	//httpDo(url,postData,"POST")

	//10 Modify courier code
	//var url string ="http://api.trackingmore.com/v2/trackings/update"
	//var postData string = "{\"tracking_number\": \"EA152563242CN\",\"carrier_code\":\"china-ems\",\"update_carrier_code\":\"china-post\"}"
	//httpDo(url,postData,"POST")

	//11 Get user info
	//var url string ="http://api.trackingmore.com/v2/trackings/getuserinfo"
	//var postData string = ""
	//httpDo(url,postData,"GET")

	//12 Get status number
	//var url string ="http://api.trackingmore.com/v2/trackings/getstatusnumber"
	//var postData string = ""
	//httpDo(url,postData,"GET")

	//13 Set number not update
	//var url string ="http://api.trackingmore.com/v2/trackings/notupdate"
	//var postData string = "[{\"tracking_number\":\"LK032051658CN\",\"carrier_code\":\"china-ems\"},{\"tracking_number\":\"EA152565249CN\",\"carrier_code\":\"china-ems\"}]"
	//httpDo(url,postData,"POST")

	//14 Get remote iterm results
	//var url string = "http://api.trackingmore.com/v2/trackings/remote"
	//var postData string = "[{\"country\":\"CN\",\"postcode\":\"400422\"},{\"country\":\"CN\",\"postcode\":\"412000\"}]"
	//httpDo(url,postData,"POST")

	//15 Get cost time iterm results
	//var url string = "http://api.trackingmore.com/v2/trackings/costtime"
	//var postData string = "[{\"carrier_code\":\"dhl\",\"destination\":\"US\",\"original\":\"CN\"},{\"carrier_code\":\"dhl\",\"destination\":\"RU\",\"original\":\"CN\"}]"
	//httpDo(url,postData,"POST")

	//16 Delete multiple tracking item
	//var url string = "http://api.trackingmore.com/v2/trackings/delete"
	//var postData string = "[{\"tracking_number\":\"EA152563242CN\",\"carrier_code\":\"china-ems\"},{\"tracking_number\":\"EA152563254CN\",\"carrier_code\":\"china-ems\"}]"
	//httpDo(url,postData,"POST")

	//17 Update multiple Tracking item
	//var url string = "http://api.trackingmore.com/v2/trackings/updatemore"
	//var postData string = "[{\"tracking_number\":\"EA152563242CN\",\"carrier_code\":\"china-post\",\"title\": \"testtitle\",\"customer_name\":\"python test\",\"customer_email\":\"942632688@qq.com\",\"order_id\":\"#123\",\"logistics_channel\":\"chase chen\",\"destination_code\":\"US\",\"status\":\"7\"},{\"tracking_number\":\"EA152563246CN\",\"carrier_code\":\"china-ems\",\"title\": \"testtitle\",\"customer_name\":\"python test\",\"customer_email\":\"942632688@qq.com\",\"order_id\":\"#123\",\"logistics_channel\":\"chase chen\",\"destination_code\":\"US\",\"status\":\"7\"}]"
	//httpDo(url,postData,"POST")

}

// 4209827392748927005485000009927294
func httpDo(url, postData, method string) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, strings.NewReader(postData))
	if err != nil {
		// handle error
	}
	req.Header.Set("Content-Type", "application/json'")
	req.Header.Set("Trackingmore-Api-Key", "61a12315-3731-483c-b2c7-172b959ee170")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	fmt.Println(string(body))
}
