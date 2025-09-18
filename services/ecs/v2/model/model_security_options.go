package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SecurityOptions struct {

	// 是否支持安全启动 true:支持安全启动 false:不支持安全启动 默认值：false
	SecureBootEnabled bool `json:"secure_boot_enabled"`

	// 是否支持vtpm启动 true:支持vtpm启动 false:不支持vtpm启动 默认值：false
	TpmEnabled bool `json:"tpm_enabled"`
}

func (o SecurityOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityOptions struct{}"
	}

	return strings.Join([]string{"SecurityOptions", string(data)}, " ")
}
