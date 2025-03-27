package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateVpnLogConfigRequestBody struct {
	LogConfig *UpdateVpnLogConfigRequestBodyContent `json:"log_config"`
}

func (o UpdateVpnLogConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpnLogConfigRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateVpnLogConfigRequestBody", string(data)}, " ")
}
