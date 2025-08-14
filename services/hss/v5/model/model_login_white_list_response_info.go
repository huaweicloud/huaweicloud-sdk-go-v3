package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LoginWhiteListResponseInfo 登录白名单详情
type LoginWhiteListResponseInfo struct {

	// **参数解释**： 服务器私有IP **取值范围**： 字符长度1-128位
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**： 登录源IP **取值范围**： 字符长度1-256位
	LoginIp *string `json:"login_ip,omitempty"`

	// **参数解释**： 登录用户名 **取值范围**： 字符长度1-256位
	LoginUserName *string `json:"login_user_name,omitempty"`

	// 更新时间，毫秒
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 备注
	Remarks *string `json:"remarks,omitempty"`

	// 企业项目名称
	EnterpriseProjectName *string `json:"enterprise_project_name,omitempty"`
}

func (o LoginWhiteListResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoginWhiteListResponseInfo struct{}"
	}

	return strings.Join([]string{"LoginWhiteListResponseInfo", string(data)}, " ")
}
