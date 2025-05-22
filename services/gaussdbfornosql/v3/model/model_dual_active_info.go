package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DualActiveInfo struct {

	// 双活角色。
	Role *string `json:"role,omitempty"`

	// 双活状态。
	Status *string `json:"status,omitempty"`

	// 双活对端实例id。
	DestinationInstanceId *string `json:"destination_instance_id,omitempty"`

	// 双活对端region。
	DestinationRegion *string `json:"destination_region,omitempty"`
}

func (o DualActiveInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DualActiveInfo struct{}"
	}

	return strings.Join([]string{"DualActiveInfo", string(data)}, " ")
}
