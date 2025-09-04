package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowUserChargeTypeResultMainResourceList struct {

	// 状态
	Status *string `json:"status,omitempty"`

	// 服务类型
	ServiceType *string `json:"service_type,omitempty"`

	// 资源类型
	ResourceType *string `json:"resource_type,omitempty"`

	// 资源包类型
	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`
}

func (o ShowUserChargeTypeResultMainResourceList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserChargeTypeResultMainResourceList struct{}"
	}

	return strings.Join([]string{"ShowUserChargeTypeResultMainResourceList", string(data)}, " ")
}
