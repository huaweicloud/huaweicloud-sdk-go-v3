package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MemberStatus 后端服务器操作状态。
type MemberStatus struct {

	// 参数解释：监听器ID
	ListenerId string `json:"listener_id"`

	// 参数解释：后端服务器的健康状态。  取值范围： - ONLINE：后端服务器正常。 - NO_MONITOR：后端服务器所在的服务器组没有健康检查器。 - OFFLINE：后端服务器关联的ECS服务器不存在或已关机。
	OperatingStatus string `json:"operating_status"`
}

func (o MemberStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MemberStatus struct{}"
	}

	return strings.Join([]string{"MemberStatus", string(data)}, " ")
}
