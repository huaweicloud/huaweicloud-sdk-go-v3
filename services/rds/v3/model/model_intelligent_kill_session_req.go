package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IntelligentKillSessionReq 一键自动kill会话请求体
type IntelligentKillSessionReq struct {

	// 是否自动SQL限流
	AutoAddSqlLimitRule *bool `json:"auto_add_sql_limit_rule,omitempty"`
}

func (o IntelligentKillSessionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IntelligentKillSessionReq struct{}"
	}

	return strings.Join([]string{"IntelligentKillSessionReq", string(data)}, " ")
}
