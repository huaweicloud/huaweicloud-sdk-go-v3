package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDDoSConnectionNumberRequest Request Object
type ListDDoSConnectionNumberRequest struct {

	// 开始时间（毫秒时间戳）
	StartTime string `json:"start_time"`

	// 结束时间（毫秒时间戳）
	EndTime string `json:"end_time"`

	// 实例id
	InstanceId string `json:"instance_id"`

	// 高防ip
	Ip string `json:"ip"`
}

func (o ListDDoSConnectionNumberRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDDoSConnectionNumberRequest struct{}"
	}

	return strings.Join([]string{"ListDDoSConnectionNumberRequest", string(data)}, " ")
}
