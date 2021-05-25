package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteRecordConfigRequest struct {
	// 直播播放域名

	Domain string `json:"domain"`
	// 流应用名称

	AppName string `json:"app_name"`
}

func (o DeleteRecordConfigRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteRecordConfigRequest struct{}"
	}

	return strings.Join([]string{"DeleteRecordConfigRequest", string(data)}, " ")
}
