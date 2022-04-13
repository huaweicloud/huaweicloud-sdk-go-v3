package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListL7rulesResponse struct {
	// 转发规则对象的列表

	Rules          *[]L7ruleResp `json:"rules,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListL7rulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListL7rulesResponse struct{}"
	}

	return strings.Join([]string{"ListL7rulesResponse", string(data)}, " ")
}
