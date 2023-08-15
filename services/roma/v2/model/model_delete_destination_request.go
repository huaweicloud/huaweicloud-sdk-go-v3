package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDestinationRequest Request Object
type DeleteDestinationRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 规则ID
	RuleId string `json:"rule_id"`

	// 目标数据源ID
	DestinationId int32 `json:"destination_id"`
}

func (o DeleteDestinationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDestinationRequest struct{}"
	}

	return strings.Join([]string{"DeleteDestinationRequest", string(data)}, " ")
}
