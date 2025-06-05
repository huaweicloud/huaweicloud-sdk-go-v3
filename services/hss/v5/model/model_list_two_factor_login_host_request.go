package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTwoFactorLoginHostRequest Request Object
type ListTwoFactorLoginHostRequest struct {

	// **参数解释**: 主机所属的企业项目ID。 企业项目ID默认取值为“0”，表示默认企业项目。如果需要查询所有企业项目下的主机，请传参“all_granted_eps”。如果您只有某个企业项目的权限，则需要传递该企业项目ID，查询该企业项目下的主机，否则会因权限不足而报错。 **约束限制**: 开通企业项目功能后才需要配置企业项目。 **取值范围**: 字符长度1-256位 **默认取值**: 0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 主机名
	HostName *string `json:"host_name,omitempty"`

	// 显示名称
	DisplayName *string `json:"display_name,omitempty"`

	// 默认10
	Limit int32 `json:"limit"`

	// 默认0
	Offset int32 `json:"offset"`
}

func (o ListTwoFactorLoginHostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTwoFactorLoginHostRequest struct{}"
	}

	return strings.Join([]string{"ListTwoFactorLoginHostRequest", string(data)}, " ")
}
