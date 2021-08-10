package model

import (
	"encoding/json"

	"strings"
)

type ServiceInfo struct {
	// 开通该子服务时间。

	CreateTime string `json:"create_time"`
	// 是否开通该子服务。

	SubscribeStatus bool `json:"subscribe_status"`
}

func (o ServiceInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ServiceInfo struct{}"
	}

	return strings.Join([]string{"ServiceInfo", string(data)}, " ")
}
