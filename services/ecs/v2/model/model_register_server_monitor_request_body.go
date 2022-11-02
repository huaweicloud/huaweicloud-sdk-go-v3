package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type RegisterServerMonitorRequestBody struct {

	// 注册云服务器监控。
	MonitorMetrics string `json:"monitorMetrics"`
}

func (o RegisterServerMonitorRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterServerMonitorRequestBody struct{}"
	}

	return strings.Join([]string{"RegisterServerMonitorRequestBody", string(data)}, " ")
}
