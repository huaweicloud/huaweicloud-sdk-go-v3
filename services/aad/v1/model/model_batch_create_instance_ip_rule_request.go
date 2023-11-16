package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateInstanceIpRuleRequest Request Object
type BatchCreateInstanceIpRuleRequest struct {

	// 实例Id
	InstanceId string `json:"instance_id"`

	// 单个 IP
	Ip string `json:"ip"`

	Body *BatchTransferRuleBody `json:"body,omitempty"`
}

func (o BatchCreateInstanceIpRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateInstanceIpRuleRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateInstanceIpRuleRequest", string(data)}, " ")
}
