package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type ScopeTokenResult struct {

	// 获取token的方式。
	Methods []string `json:"methods" xml:"methods"`

	// token过期时间。
	ExpiresAt string `json:"expires_at" xml:"expires_at"`

	// 服务目录信息。
	Catalog *[]TokenCatalog `json:"catalog,omitempty" xml:"catalog"`

	Domain *TokenDomainResult `json:"domain" xml:"domain"`

	Project *TokenProjectResult `json:"project" xml:"project"`

	// token的权限信息。
	Roles []TokenRole `json:"roles" xml:"roles"`

	User *ScopedTokenUser `json:"user" xml:"user"`

	// token下发时间。
	IssuedAt string `json:"issued_at" xml:"issued_at"`
}

func (o ScopeTokenResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScopeTokenResult struct{}"
	}

	return strings.Join([]string{"ScopeTokenResult", string(data)}, " ")
}
