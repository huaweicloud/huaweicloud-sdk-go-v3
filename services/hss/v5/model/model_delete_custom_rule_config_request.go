package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCustomRuleConfigRequest Request Object
type DeleteCustomRuleConfigRequest struct {
	Body *DeleteCustomRuleIdsRequestInfo `json:"body,omitempty"`
}

func (o DeleteCustomRuleConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCustomRuleConfigRequest struct{}"
	}

	return strings.Join([]string{"DeleteCustomRuleConfigRequest", string(data)}, " ")
}
