package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListenerMemberInfo 后端服务器监听器粒度的健康检查结果
type ListenerMemberInfo struct {

	// 后端服务器关联的监听器id。
	ListenerId string `json:"listener_id"`

	// 后端服务器的健康状态。  取值： - ONLINE：后端服务器正常。 - NO_MONITOR：后端服务器所在的服务器组没有健康检查器。 - OFFLINE：后端服务器关联的ECS服务器不存在或已关机或服务异常。
	OperatingStatus string `json:"operating_status"`
}

func (o ListenerMemberInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListenerMemberInfo struct{}"
	}

	return strings.Join([]string{"ListenerMemberInfo", string(data)}, " ")
}
