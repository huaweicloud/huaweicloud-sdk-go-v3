package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTaskRulesetResponse Response Object
type ListTaskRulesetResponse struct {
	Body           *[]TaskRulesetInfo `json:"body,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListTaskRulesetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTaskRulesetResponse struct{}"
	}

	return strings.Join([]string{"ListTaskRulesetResponse", string(data)}, " ")
}
