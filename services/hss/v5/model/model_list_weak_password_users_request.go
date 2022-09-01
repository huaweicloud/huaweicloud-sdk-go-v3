package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListWeakPasswordUsersRequest struct {

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty" xml:"enterprise_project_id"`

	// 服务器名称
	HostName *string `json:"host_name,omitempty" xml:"host_name"`

	// 服务器IP地址
	HostIp *string `json:"host_ip,omitempty" xml:"host_ip"`

	// 弱口令账号名称
	UserName *string `json:"user_name,omitempty" xml:"user_name"`

	// 主机ID，不赋值时，查租户所有主机
	HostId *string `json:"host_id,omitempty" xml:"host_id"`

	// 每页数量
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset *int32 `json:"offset,omitempty" xml:"offset"`
}

func (o ListWeakPasswordUsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWeakPasswordUsersRequest struct{}"
	}

	return strings.Join([]string{"ListWeakPasswordUsersRequest", string(data)}, " ")
}
