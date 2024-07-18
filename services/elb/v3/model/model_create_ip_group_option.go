package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIpGroupOption 创建IP地址组请求参数。
type CreateIpGroupOption struct {

	// 参数解释：IP地址组所在的项目ID。
	ProjectId *string `json:"project_id,omitempty"`

	// 参数解释：IP地址组的描述。
	Description *string `json:"description,omitempty"`

	// 参数解释：IP地址组的名称。
	Name *string `json:"name,omitempty"`

	// 参数解释：IP地址组中包含的IP或网段列表。[]表示任意IP。
	IpList []CreateIpGroupIpOption `json:"ip_list"`

	// 参数解释：IP地址组所在的企业项目ID。  [不支持该字段，请勿使用。](tag:dt,dt_test,hcso_dt)
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o CreateIpGroupOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIpGroupOption struct{}"
	}

	return strings.Join([]string{"CreateIpGroupOption", string(data)}, " ")
}
