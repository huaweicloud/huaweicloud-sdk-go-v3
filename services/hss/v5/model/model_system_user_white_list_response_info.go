package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SystemUserWhiteListResponseInfo 系统用户白名单详情
type SystemUserWhiteListResponseInfo struct {

	// 企业项目名称
	EnterpriseProjectName *string `json:"enterprise_project_name,omitempty"`

	// 服务器ID
	HostId *string `json:"host_id,omitempty"`

	// 服务器名称
	HostName *string `json:"host_name,omitempty"`

	// 服务器私有IP
	PrivateIp *string `json:"private_ip,omitempty"`

	// 弹性公网IP地址
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
