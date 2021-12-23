package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建IP地址组请求参数。
type CreateIpGroupOption struct {
	// IP地址组所在的项目ID。

	ProjectId *string `json:"project_id,omitempty"`
	// IP地址组的描述。

	Description *string `json:"description,omitempty"`
	// IP地址组的名称。

	Name *string `json:"name,omitempty"`
	// IP地址组中包含的IP或网段列表。[]表示任意IP。

	IpList []CreateIpGroupIpOption `json:"ip_list"`
	// IP地址组所在的企业项目ID。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o CreateIpGroupOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIpGroupOption struct{}"
	}

	return strings.Join([]string{"CreateIpGroupOption", string(data)}, " ")
}
