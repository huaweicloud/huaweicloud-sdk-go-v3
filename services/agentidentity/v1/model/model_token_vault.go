package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TokenVault struct {

	// The unique identifier of the token vault.
	TokenVaultId string `json:"token_vault_id"`

	// TokenVault 对象统一资源标识（URN）。
	Urn string `json:"urn"`

	KmsConfiguration *KmsConfiguration `json:"kms_configuration"`

	PolicyEngineConfiguration *PolicyEngineConfiguration `json:"policy_engine_configuration,omitempty"`

	// Timestamp in RFC 3339 format (UTC)
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	// 自定义标签列表。
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o TokenVault) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TokenVault struct{}"
	}

	return strings.Join([]string{"TokenVault", string(data)}, " ")
}
