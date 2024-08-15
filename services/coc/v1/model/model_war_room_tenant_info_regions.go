package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WarRoomTenantInfoRegions struct {

	// 主键
	Code *string `json:"code,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`
}

func (o WarRoomTenantInfoRegions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WarRoomTenantInfoRegions struct{}"
	}

	return strings.Join([]string{"WarRoomTenantInfoRegions", string(data)}, " ")
}
