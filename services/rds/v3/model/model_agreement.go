package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Agreement struct {

	// 许可ID。
	Id *string `json:"id,omitempty"`

	// 许可名称。
	Name *string `json:"name,omitempty"`

	// 许可语言类型。
	Language *string `json:"language,omitempty"`

	// 许可版本。
	Version *string `json:"version,omitempty"`

	// 许可链接。
	ProvisionUrl *string `json:"provision_url,omitempty"`
}

func (o Agreement) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Agreement struct{}"
	}

	return strings.Join([]string{"Agreement", string(data)}, " ")
}
