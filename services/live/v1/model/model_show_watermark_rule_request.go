package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWatermarkRuleRequest Request Object
type ShowWatermarkRuleRequest struct {

	// 规则ID
	Id string `json:"id"`
}

func (o ShowWatermarkRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWatermarkRuleRequest struct{}"
	}

	return strings.Join([]string{"ShowWatermarkRuleRequest", string(data)}, " ")
}
