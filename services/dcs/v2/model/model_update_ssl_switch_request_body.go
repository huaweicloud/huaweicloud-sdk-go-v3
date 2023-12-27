package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateSslSwitchRequestBody struct {

	// 开启或关闭SSL。true：开启/false：关闭
	Enabled bool `json:"enabled"`
}

func (o UpdateSslSwitchRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSslSwitchRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateSslSwitchRequestBody", string(data)}, " ")
}
