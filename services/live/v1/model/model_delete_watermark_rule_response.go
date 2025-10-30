package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteWatermarkRuleResponse Response Object
type DeleteWatermarkRuleResponse struct {
	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteWatermarkRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWatermarkRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteWatermarkRuleResponse", string(data)}, " ")
}
