package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceIpRuleRequest Request Object
type ListInstanceIpRuleRequest struct {

	// 实例Id
	InstanceId string `json:"instance_id"`

	// 单个 IP
	Ip string `json:"ip"`
}

func (o ListInstanceIpRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceIpRuleRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceIpRuleRequest", string(data)}, " ")
}
