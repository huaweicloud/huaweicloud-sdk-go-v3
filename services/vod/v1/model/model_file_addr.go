package model

import (
	"encoding/json"

	"strings"
)

type FileAddr struct {
	// OBS的bucket名称。<br/>

	Bucket string `json:"bucket"`
	// 输入OBS Bucket所在数据中心（OBS Location）。<br/>

	Location string `json:"location"`
	// OBS对象路径，遵守OSS Object定义。<br/> 当用于指示input时,需要指定到具体对象<br/> 当用于指示output时, 只需指定到转码结果期望存放的路径<br/>

	Object string `json:"object"`
}

func (o FileAddr) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "FileAddr struct{}"
	}

	return strings.Join([]string{"FileAddr", string(data)}, " ")
}
