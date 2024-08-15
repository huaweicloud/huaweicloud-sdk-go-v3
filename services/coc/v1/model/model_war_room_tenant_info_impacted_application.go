package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WarRoomTenantInfoImpactedApplication struct {

	// 主键
	Id *string `json:"id,omitempty"`

	// 名字
	Name *string `json:"name,omitempty"`
}

func (o WarRoomTenantInfoImpactedApplication) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WarRoomTenantInfoImpactedApplication struct{}"
	}

	return strings.Join([]string{"WarRoomTenantInfoImpactedApplication", string(data)}, " ")
}
