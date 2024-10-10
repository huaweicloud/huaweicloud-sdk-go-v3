package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDDoSFlowRequest Request Object
type ListDDoSFlowRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 高防IP
	Ip string `json:"ip"`

	// 请求类型 pps、bps
	Type string `json:"type"`

	// 开始时间（毫秒时间戳）
	StartTime string `json:"start_time"`

	// 结束时间（毫秒时间戳）
	EndTime string `json:"end_time"`
}

func (o ListDDoSFlowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDDoSFlowRequest struct{}"
	}

	return strings.Join([]string{"ListDDoSFlowRequest", string(data)}, " ")
}
