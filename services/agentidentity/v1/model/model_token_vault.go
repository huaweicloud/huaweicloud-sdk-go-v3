package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TokenVault struct {

	// The unique identifier of the token vault.
	TokenVaultId string `json:"token_vault_id"`

	KmsConfiguration *KmsConfiguration `json:"kms_configuration"`

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
