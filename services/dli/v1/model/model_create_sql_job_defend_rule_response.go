package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSqlJobDefendRuleResponse Response Object
type CreateSqlJobDefendRuleResponse struct {

	// 执行请求是否成功。“true”表示请求执行成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 系统提示信息，执行成功时，信息可能为空。
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateSqlJobDefendRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSqlJobDefendRuleResponse struct{}"
	}

	return strings.Join([]string{"CreateSqlJobDefendRuleResponse", string(data)}, " ")
}
