package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/resty.v1"
	"time"
)

func main() {
	url := fmt.Sprintf("https://status.sto.shopee.io/pubapi/azmetainfo/nodelist?role=%s&ips=%s",
		"unbound", "10.66.243.130")

	response, _ := resty.New().SetTimeout(time.Minute).R().SetHeaders(map[string]string{
		"token": "f9e644cc57b52f7604741255f52bb9024e7e7923",
	}).Get(url)
	println(response.String())
	print(response.StatusCode())
	type respStruct struct {
		Nodes []struct {
			ID string `json:"id"`
			IP string `json:"ip_lan"`

			Role        string `json:"role"`
			Status      string `json:"status"`
			ClusterID   string `json:"cluster_id"`
			ClusterName string `json:"cluster_name"`
		} `json:"nodes"`
	}
	var respObj respStruct
	if err := json.Unmarshal([]byte(response.String()), &respObj); err != nil {

		println("error1 ")
	}

	if len(respObj.Nodes) == 0 {

		print("error2")
	}

	NodeID := respObj.Nodes[0].ID
	println(response.String())
	println(response.StatusCode())
	println(NodeID)

}
