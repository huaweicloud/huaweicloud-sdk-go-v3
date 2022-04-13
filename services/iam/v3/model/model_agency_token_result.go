package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type AgencyTokenResult struct {
	// 获取token的方式。

	Methods []string `json:"methods"`
	// token到期时间。

	ExpiresAt string `json:"expires_at"`
	// token下发时间。

	IssuedAt string `json:"issued_at"`

	AssumedBy *AgencyAssumedby `json:"assumed_by"`
	// 服务目录信息。

	Catalog *[]TokenCatalog `json:"catalog,omitempty"`

	Domain *AgencyTokenDomain `json:"domain,omitempty"`

	Project *AgencyTokenProject `json:"project,omitempty"`
	// 委托token的权限信息。

	Roles []TokenRole `json:"roles"`

	User *AgencyTokenUser `json:"user"`
}

func (o AgencyTokenResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgencyTokenResult struct{}"
	}

	return strings.Join([]string{"AgencyTokenResult", string(data)}, " ")
}
