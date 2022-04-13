package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type NovaDeleteServerRequest struct {
	// 云服务器ID。

	ServerId string `json:"server_id"`
}

func (o NovaDeleteServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaDeleteServerRequest struct{}"
	}

	return strings.Join([]string{"NovaDeleteServerRequest", string(data)}, " ")
}
