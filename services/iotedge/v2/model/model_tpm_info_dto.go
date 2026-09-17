package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TpmInfoDto TPM芯片信息
type TpmInfoDto struct {

	// 厂商信息
	ManufactureId *string `json:"manufacture_id,omitempty"`

	// 协议版本
	SpecVersion *string `json:"spec_version,omitempty"`
}

func (o TpmInfoDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TpmInfoDto struct{}"
	}

	return strings.Join([]string{"TpmInfoDto", string(data)}, " ")
}
