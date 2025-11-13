package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateOcspSwitchRequestBody struct {

	// 启用或禁用OCSP。
	OcspSwitch bool `json:"ocsp_switch"`
}

func (o UpdateOcspSwitchRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateOcspSwitchRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateOcspSwitchRequestBody", string(data)}, " ")
}
