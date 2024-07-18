package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IpGroup IP地址组信息。
type IpGroup struct {

	// 参数解释：IP地址组的ID。
	Id string `json:"id"`

	// 参数解释：IP地址组的名称。
	Name string `json:"name"`

	// 参数解释：IP地址组的描述信息。
	Description string `json:"description"`

	// 参数解释：IP地址组中包含的IP或网段列表。[]表示任意IP。
	IpList []IpInfo `json:"ip_list"`

	// 参数解释：与IP地址组关联的监听器的ID列表。
	Listeners []ListenerRef `json:"listeners"`

	// 参数解释：IP地址组的项目ID。
	ProjectId string `json:"project_id"`

	// 参数解释：IP地址组所在的企业项目ID。  [不支持该字段，请勿使用。](tag:dt,dt_test,hcso_dt)
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 参数解释：IP地址组的创建时间。
	CreatedAt string `json:"created_at"`

	// 参数解释：IP地址组的更新时间。
	UpdatedAt string `json:"updated_at"`
}

func (o IpGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpGroup struct{}"
	}

	return strings.Join([]string{"IpGroup", string(data)}, " ")
}
