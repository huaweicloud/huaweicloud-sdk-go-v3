package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RuntimeInfoDto 边缘软件运行时信息
type RuntimeInfoDto struct {

	// 是否启用TPM
	EnableTpm *bool `json:"enable_tpm,omitempty"`
}

func (o RuntimeInfoDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuntimeInfoDto struct{}"
	}

	return strings.Join([]string{"RuntimeInfoDto", string(data)}, " ")
}
