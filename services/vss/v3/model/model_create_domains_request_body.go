package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateDomainsRequestBody struct {
	// 域名

	DomainName string `json:"domain_name"`
	// 域名的别名

	Alias string `json:"alias"`
}

func (o CreateDomainsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDomainsRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDomainsRequestBody", string(data)}, " ")
}
