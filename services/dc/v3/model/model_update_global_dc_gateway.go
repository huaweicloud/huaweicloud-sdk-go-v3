package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGlobalDcGateway 更新global-dc-gateway的属性详情。
type UpdateGlobalDcGateway struct {

	// global-dc-gateway名字。
	Name *string `json:"name,omitempty"`

	// 描述信息。
	Description *string `json:"description,omitempty"`

	// global-dc-gateway的地址族
	AddressFamily *string `json:"address_family,omitempty"`
}

func (o UpdateGlobalDcGateway) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGlobalDcGateway struct{}"
	}

	return strings.Join([]string{"UpdateGlobalDcGateway", string(data)}, " ")
}
