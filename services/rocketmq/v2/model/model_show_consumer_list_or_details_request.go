package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowConsumerListOrDetailsRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 消费组名称。
	Group string `json:"group" xml:"group"`

	// 待查询的topic，不指定时查询topic列表，指定时查询详情。
	Topic *string `json:"topic,omitempty" xml:"topic"`
}

func (o ShowConsumerListOrDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConsumerListOrDetailsRequest struct{}"
	}

	return strings.Join([]string{"ShowConsumerListOrDetailsRequest", string(data)}, " ")
}
