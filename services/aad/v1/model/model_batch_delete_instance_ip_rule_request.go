package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteInstanceIpRuleRequest Request Object
type BatchDeleteInstanceIpRuleRequest struct {

	// 实例Id
	InstanceId string `json:"instance_id"`

	// 单个 IP
	Ip string `json:"ip"`

	Body *BatchIdBody `json:"body,omitempty"`
}

func (o BatchDeleteInstanceIpRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteInstanceIpRuleRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteInstanceIpRuleRequest", string(data)}, " ")
}
