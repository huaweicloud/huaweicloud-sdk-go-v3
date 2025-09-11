package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RuleSqlIdRequest struct {

	// SQL规则ID
	Id string `json:"id"`

	// 状态  - OFF：关闭  - ON：开启
	Status string `json:"status"`
}

func (o RuleSqlIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleSqlIdRequest struct{}"
	}

	return strings.Join([]string{"RuleSqlIdRequest", string(data)}, " ")
}
