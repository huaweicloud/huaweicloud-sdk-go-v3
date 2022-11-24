package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceBasicInfo struct {

	// 资源类型编码。例如：hws.resource.type.general。
	ResourceTypeCode *string `json:"resource_type_code,omitempty"`

	// 资源类型归属的服务类型编码。例如：hws.service.type.offline。
	ProductOwnerService *string `json:"product_owner_service,omitempty"`

	// 资源类型名称。例如：通用规格。
	Name *string `json:"name,omitempty"`

	// 资源类型描述。
	Description *string `json:"description,omitempty"`
}

func (o ResourceBasicInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceBasicInfo struct{}"
	}

	return strings.Join([]string{"ResourceBasicInfo", string(data)}, " ")
}
