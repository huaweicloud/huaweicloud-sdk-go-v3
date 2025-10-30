package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWatermarkRuleRequest Request Object
type UpdateWatermarkRuleRequest struct {

	// 规则ID，在创建成功后返回
	Id string `json:"id"`

	Body *ModifyWatermarkRule `json:"body,omitempty"`
}

func (o UpdateWatermarkRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWatermarkRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateWatermarkRuleRequest", string(data)}, " ")
}
