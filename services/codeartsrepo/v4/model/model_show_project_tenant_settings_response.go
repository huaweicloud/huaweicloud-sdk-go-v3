package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProjectTenantSettingsResponse Response Object
type ShowProjectTenantSettingsResponse struct {

	// **参数解释：** 仓库默认加密设置是否开启。
	DefaultEncryptionEnabled *bool `json:"default_encryption_enabled,omitempty"`

	// **参数解释：** 租户设置的加密类型。 **取值范围：** KMS,normal,null,当为KMS时表示开启了KMS加密。
	EncryptionType *string `json:"encryption_type,omitempty"`

	// **参数解释：** 允许公共访问。 **取值范围：** allow 允许 deny 拒绝。
	PermitPublic   *string `json:"permit_public,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowProjectTenantSettingsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectTenantSettingsResponse struct{}"
	}

	return strings.Join([]string{"ShowProjectTenantSettingsResponse", string(data)}, " ")
}
