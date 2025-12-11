package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SystemUserWhiteListResponseInfo 系统用户白名单详情
type SystemUserWhiteListResponseInfo struct {

	// 企业项目名称
	EnterpriseProjectName *string `json:"enterprise_project_name,omitempty"`

	// **参数解释**： 服务器（主机）的唯一标识ID **取值范围**： 字符长度1-64位
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 服务器名称 **取值范围**: 字符长度1-256位
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**： 服务器私有IP **取值范围**： 字符长度1-128位
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**： 弹性公网IP地址 **取值范围**： 字符长度1-256位，支持IPv4或IPv6格式（IPv4长度7-15位，IPv6长度15-39位）
	PublicIp *string `json:"public_ip,omitempty"`

	// 系统用户名列表
	SystemUserNameList *[]string `json:"system_user_name_list,omitempty"`

	// 更新时间，毫秒
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 备注
	Remarks *string `json:"remarks,omitempty"`
}

func (o SystemUserWhiteListResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SystemUserWhiteListResponseInfo struct{}"
	}

	return strings.Join([]string{"SystemUserWhiteListResponseInfo", string(data)}, " ")
}
