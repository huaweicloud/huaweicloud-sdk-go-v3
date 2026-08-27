package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VerifyProviderResponse Response Object
type VerifyProviderResponse struct {

	// 供应商id。
	ProviderId *string `json:"provider_id,omitempty"`

	// 是否验证成功。
	Success *bool `json:"success,omitempty"`

	// 验证结果信息。
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o VerifyProviderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VerifyProviderResponse struct{}"
	}

	return strings.Join([]string{"VerifyProviderResponse", string(data)}, " ")
}
