package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteWhiteBlackIpRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteWhiteBlackIpRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWhiteBlackIpRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteWhiteBlackIpRuleResponse", string(data)}, " ")
}
