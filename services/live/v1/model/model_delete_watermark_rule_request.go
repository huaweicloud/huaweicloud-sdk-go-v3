package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteWatermarkRuleRequest Request Object
type DeleteWatermarkRuleRequest struct {

	// 规则ID
	Id string `json:"id"`
}

func (o DeleteWatermarkRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWatermarkRuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteWatermarkRuleRequest", string(data)}, " ")
}
