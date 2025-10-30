package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWatermarkRuleRequest Request Object
type CreateWatermarkRuleRequest struct {
	Body *WatermarkRule `json:"body,omitempty"`
}

func (o CreateWatermarkRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWatermarkRuleRequest struct{}"
	}

	return strings.Join([]string{"CreateWatermarkRuleRequest", string(data)}, " ")
}
