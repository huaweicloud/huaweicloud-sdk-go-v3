package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteL7RuleResponse Response Object
type DeleteL7RuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteL7RuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteL7RuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteL7RuleResponse", string(data)}, " ")
}
