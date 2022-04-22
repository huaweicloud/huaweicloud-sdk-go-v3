package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新网络实例的详细信息。
type UpdateNetworkInstance struct {

	// 网络实例的名字。
	Name *string `json:"name,omitempty"`

	// 网络实例的描述。
	Description *string `json:"description,omitempty"`

	// VPC或者VGW发布的网段路由列表，ER场景不需要此字段。
	Cidrs *[]string `json:"cidrs,omitempty"`
}

func (o UpdateNetworkInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNetworkInstance struct{}"
	}

	return strings.Join([]string{"UpdateNetworkInstance", string(data)}, " ")
}
