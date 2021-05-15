package json_demo

import (
	"encoding/json"
	"fmt"
	"log"
)

type MovieResp struct {
	Code int
	Message string `json:"message"` // 成员标签 filed tag （json key 别名）
	Data []string `json:"dataArray,omitempty"`  // 成员标签 filed tag （omitempty为空则不转json）
}

func TestJson()  {
	resp := MovieResp{
		Code: 0,
		Message: "success",
		Data: []string{"aa", "bb"},
	}
	jsonStr1 , err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", jsonStr1)

	jsonStr2 , err := json.MarshalIndent(resp, "", "	")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", jsonStr2)

	resp1 := MovieResp{}
	err = json.Unmarshal(jsonStr1, &resp1)
	if err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(resp1)
}
