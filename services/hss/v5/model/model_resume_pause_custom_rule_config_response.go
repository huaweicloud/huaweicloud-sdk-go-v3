package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResumePauseCustomRuleConfigResponse Response Object
type ResumePauseCustomRuleConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ResumePauseCustomRuleConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResumePauseCustomRuleConfigResponse struct{}"
	}

	return strings.Join([]string{"ResumePauseCustomRuleConfigResponse", string(data)}, " ")
}
