package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableDnsRequest Request Object
type EnableDnsRequest struct {

	// 实例ID，从[查询所有实例列表](ListInstancesDetails.xml)获取实例ID。
	InstanceId string `json:"instance_id"`
}

func (o EnableDnsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableDnsRequest struct{}"
	}

	return strings.Join([]string{"EnableDnsRequest", string(data)}, " ")
}
