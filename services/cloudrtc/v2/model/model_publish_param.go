package model

import (
	"encoding/json"

	"strings"
)

// 转推参数
type PublishParam struct {
	// 合流任务完成后，转推的RTMP推流地址。

	RtmpUrls []string `json:"rtmp_urls"`
}

func (o PublishParam) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "PublishParam struct{}"
	}

	return strings.Join([]string{"PublishParam", string(data)}, " ")
}
