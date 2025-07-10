package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRuleNewResponse Response Object
type DeleteRuleNewResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteRuleNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRuleNewResponse struct{}"
	}

	return strings.Join([]string{"DeleteRuleNewResponse", string(data)}, " ")
}
