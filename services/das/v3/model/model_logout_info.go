package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogoutInfo 登出信息
type LogoutInfo struct {

	// 登出数据库的用户名
	Username *string `json:"username,omitempty"`
}

func (o LogoutInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogoutInfo struct{}"
	}

	return strings.Join([]string{"LogoutInfo", string(data)}, " ")
}
