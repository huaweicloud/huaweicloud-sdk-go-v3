package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GpgKeyDto struct {

	// gpg名称
	GpgName *string `json:"gpg_name,omitempty"`

	// 是否开启gpg认证
	OpenGpgVerified *bool `json:"open_gpg_verified,omitempty"`

	// gpg认证状态
	VerificationStatus *int32 `json:"verification_status,omitempty"`

	// gpg初始化id
	GpgPrimaryKeyId *string `json:"gpg_primary_key_id,omitempty"`

	// gpg昵称
	GpgNickName *string `json:"gpg_nick_name,omitempty"`

	// gpg租户昵称
	GpgTenantName *string `json:"gpg_tenant_name,omitempty"`

	// gpg用户名称
	GpgUserName *string `json:"gpg_user_name,omitempty"`
}

func (o GpgKeyDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GpgKeyDto struct{}"
	}

	return strings.Join([]string{"GpgKeyDto", string(data)}, " ")
}
