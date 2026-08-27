package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModelGroupProviderSimpleResp 模型组详情中供应商的简化信息（仅id和name）。
type ModelGroupProviderSimpleResp struct {

	// 供应商id。
	Id *string `json:"id,omitempty"`

	// 供应商名称。
	ProviderName *string `json:"provider_name,omitempty"`
}

func (o ModelGroupProviderSimpleResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModelGroupProviderSimpleResp struct{}"
	}

	return strings.Join([]string{"ModelGroupProviderSimpleResp", string(data)}, " ")
}
