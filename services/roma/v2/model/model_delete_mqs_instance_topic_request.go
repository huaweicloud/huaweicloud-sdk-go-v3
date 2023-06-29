package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteMqsInstanceTopicRequest Request Object
type DeleteMqsInstanceTopicRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 待删除的topic，多个topic以“,”分割。
	Name string `json:"name"`
}

func (o DeleteMqsInstanceTopicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMqsInstanceTopicRequest struct{}"
	}

	return strings.Join([]string{"DeleteMqsInstanceTopicRequest", string(data)}, " ")
}
