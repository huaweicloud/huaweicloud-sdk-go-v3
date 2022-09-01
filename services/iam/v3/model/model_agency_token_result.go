package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type AgencyTokenResult struct {

	// 获取token的方式。
	Methods []string `json:"methods" xml:"methods"`

	// token到期时间。
	ExpiresAt string `json:"expires_at" xml:"expires_at"`

	// token下发时间。
	IssuedAt string `json:"issued_at" xml:"issued_at"`

	AssumedBy *AgencyAssumedby `json:"assumed_by" xml:"assumed_by"`

	// 服务目录信息。
	Catalog *[]TokenCatalog `json:"catalog,omitempty" xml:"catalog"`

	Domain *AgencyTokenDomain `json:"domain,omitempty" xml:"domain"`

	Project *AgencyTokenProject `json:"project,omitempty" xml:"project"`

	// 委托token的权限信息。
	Roles []TokenRole `json:"roles" xml:"roles"`

	User *AgencyTokenUser `json:"user" xml:"user"`
}

func (o AgencyTokenResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgencyTokenResult struct{}"
	}

	return strings.Join([]string{"AgencyTokenResult", string(data)}, " ")
}
