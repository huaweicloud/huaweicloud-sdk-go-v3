package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddLoginWhiteListRequestInfo 登录白名单
type AddLoginWhiteListRequestInfo struct {

	// 服务器私有IP
	PrivateIp string `json:"private_ip"`

	// 登录源IP
	LoginIp string `json:"login_ip"`

	// 登录用户名
	LoginUserName string `json:"login_user_name"`

	// 备注
	Remarks *string `json:"remarks,omitempty"`

	// 是否同时处理相关告警事件
	HandleEvent *bool `json:"handle_event,omitempty"`
}

func (o AddLoginWhiteListRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddLoginWhiteListRequestInfo struct{}"
	}

	return strings.Join([]string{"AddLoginWhiteListRequestInfo", string(data)}, " ")
}
