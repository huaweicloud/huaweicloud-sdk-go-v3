package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListTaskRulesetResponse struct {
	Body           *[]ListTaskRulesetRes `json:"body,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListTaskRulesetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTaskRulesetResponse struct{}"
	}

	return strings.Join([]string{"ListTaskRulesetResponse", string(data)}, " ")
}
