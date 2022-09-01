package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListTaskRulesetResponse struct {
	Body           *[]ListTaskRulesetRes `json:"body,omitempty" xml:"body"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListTaskRulesetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTaskRulesetResponse struct{}"
	}

	return strings.Join([]string{"ListTaskRulesetResponse", string(data)}, " ")
}
