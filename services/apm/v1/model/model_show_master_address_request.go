package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMasterAddressRequest Request Object
type ShowMasterAddressRequest struct {

	// region英文名称。
	RegionName string `json:"region_name"`
}

func (o ShowMasterAddressRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMasterAddressRequest struct{}"
	}

	return strings.Join([]string{"ShowMasterAddressRequest", string(data)}, " ")
}
