package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type TokenResult struct {

	// 服务目录信息。
	Catalog []TokenCatalog `json:"catalog" xml:"catalog"`

	Domain *TokenDomainResult `json:"domain,omitempty" xml:"domain"`

	// token过期时间。
	ExpiresAt string `json:"expires_at" xml:"expires_at"`

	// token下发时间。
	IssuedAt string `json:"issued_at" xml:"issued_at"`

	// 获取token的方式。
	Methods []string `json:"methods" xml:"methods"`

	Project *TokenProjectResult `json:"project,omitempty" xml:"project"`

	// token的权限信息。
	Roles []TokenRole `json:"roles" xml:"roles"`

	User *TokenUserResult `json:"user" xml:"user"`
}

func (o TokenResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TokenResult struct{}"
	}

	return strings.Join([]string{"TokenResult", string(data)}, " ")
}
