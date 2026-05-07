package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResumePauseCustomRuleConfigRequest Request Object
type ResumePauseCustomRuleConfigRequest struct {
	Body *ResumePauseCustomRuleIdsRequestInfo `json:"body,omitempty"`
}

func (o ResumePauseCustomRuleConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResumePauseCustomRuleConfigRequest struct{}"
	}

	return strings.Join([]string{"ResumePauseCustomRuleConfigRequest", string(data)}, " ")
}
