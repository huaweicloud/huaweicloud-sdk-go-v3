package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNetworkInstance 更新网络实例的详细信息。
type UpdateNetworkInstance struct {

	// 网络实例的名字。只能由中文、英文字母、数字、下划线、中划线、点组成。
	Name *string `json:"name,omitempty"`

	// 网络实例的描述。不支持 <>。
	Description *string `json:"description,omitempty"`

	// 网络实例发布的网段路由列表。
	Cidrs *[]string `json:"cidrs,omitempty"`
}

func (o UpdateNetworkInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNetworkInstance struct{}"
	}

	return strings.Join([]string{"UpdateNetworkInstance", string(data)}, " ")
}
