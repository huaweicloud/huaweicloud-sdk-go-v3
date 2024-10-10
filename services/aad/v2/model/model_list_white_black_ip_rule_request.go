package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWhiteBlackIpRuleRequest Request Object
type ListWhiteBlackIpRuleRequest struct {

	// white-白名单，black-黑名单
	Type string `json:"type"`

	// instanceId
	InstanceId string `json:"instance_id"`
}

func (o ListWhiteBlackIpRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWhiteBlackIpRuleRequest struct{}"
	}

	return strings.Join([]string{"ListWhiteBlackIpRuleRequest", string(data)}, " ")
}
