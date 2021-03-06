package model

import (
	"encoding/json"

	"strings"
)

// 查询IP地址组返回对象
type IpGroup struct {
	// IP地址组的创建时间

	CreatedAt string `json:"created_at"`
	// IP地址组的更新时间。

	Description string `json:"description"`
	// IP地址组的id。

	Id string `json:"id"`
	// IP地址组中包含的ip或网段列表。[]表示任意ip。

	IpList []IpInfo `json:"ip_list"`
	// 与IP地址组关联的监听器的id列表。

	Listeners []ListenerRef `json:"listeners"`
	// IP地址组的名称。

	Name string `json:"name"`
	// IP地址组的项目id。

	ProjectId string `json:"project_id"`
	// IP地址组的更新时间。

	UpdatedAt string `json:"updated_at"`
}

func (o IpGroup) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "IpGroup struct{}"
	}

	return strings.Join([]string{"IpGroup", string(data)}, " ")
}
