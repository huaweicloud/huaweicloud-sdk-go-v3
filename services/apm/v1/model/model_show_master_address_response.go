package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowMasterAddressResponse struct {
	// region的id，英文名称，

	RegionName *string `json:"region_name,omitempty"`
	// APMmaster服务对对外暴露的地址，提供服务注册和心跳上报

	MasterAddress  *string `json:"master_address,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowMasterAddressResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMasterAddressResponse struct{}"
	}

	return strings.Join([]string{"ShowMasterAddressResponse", string(data)}, " ")
}
