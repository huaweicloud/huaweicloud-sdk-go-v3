package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IpGroup IP地址组信息。
type IpGroup struct {

	// IP地址组的创建时间。
	CreatedAt string `json:"created_at"`

	// IP地址组的描述信息。
	Description string `json:"description"`

	// IP地址组的ID。
	Id string `json:"id"`

	// IP地址组中包含的IP或网段列表。[]表示任意IP。
	IpList []IpInfo `json:"ip_list"`

	// 与IP地址组关联的监听器的ID列表。
	Listeners []ListenerRef `json:"listeners"`

	// IP地址组的名称。
	Name string `json:"name"`

	// IP地址组的项目ID。
	ProjectId string `json:"project_id"`

	// IP地址组的更新时间。
	UpdatedAt string `json:"updated_at"`
}

func (o IpGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpGroup struct{}"
	}

	return strings.Join([]string{"IpGroup", string(data)}, " ")
}
