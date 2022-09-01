package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// token详细信息。
type UnscopedTokenInfo struct {

	// 过期时间。
	ExpiresAt string `json:"expires_at" xml:"expires_at"`

	// token获取方式，联邦认证默认为mapped。
	Methods []string `json:"methods" xml:"methods"`

	// 生成时间。
	IssuedAt string `json:"issued_at" xml:"issued_at"`

	User *FederationUserBody `json:"user" xml:"user"`

	// roles信息。
	Roles *[]UnscopedTokenInfoRoles `json:"roles,omitempty" xml:"roles"`

	// catalog信息。
	Catalog *[]UnscopedTokenInfoCatalog `json:"catalog,omitempty" xml:"catalog"`
}

func (o UnscopedTokenInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnscopedTokenInfo struct{}"
	}

	return strings.Join([]string{"UnscopedTokenInfo", string(data)}, " ")
}
