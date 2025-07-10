package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRuleNewResponse Response Object
type CreateRuleNewResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateRuleNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRuleNewResponse struct{}"
	}

	return strings.Join([]string{"CreateRuleNewResponse", string(data)}, " ")
}
