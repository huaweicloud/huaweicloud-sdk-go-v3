package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RouterWithStatus struct {

	// 资源状态。
	Status *string `json:"status,omitempty" xml:"status"`

	// Router(VPC)所属VPC的ID。
	RouterId *string `json:"router_id,omitempty" xml:"router_id"`

	// Router(VPC)所在的region。
	RouterRegion *string `json:"router_region,omitempty" xml:"router_region"`
}

func (o RouterWithStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RouterWithStatus struct{}"
	}

	return strings.Join([]string{"RouterWithStatus", string(data)}, " ")
}
