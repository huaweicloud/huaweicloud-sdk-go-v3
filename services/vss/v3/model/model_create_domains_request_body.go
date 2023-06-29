package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateDomainsRequestBody struct {

	// 网站域名
	DomainName string `json:"domain_name"`

	// 网站域名的别名
	Alias string `json:"alias"`
}

func (o CreateDomainsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDomainsRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDomainsRequestBody", string(data)}, " ")
}
