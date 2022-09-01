package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateProjectDomainRequestBody struct {

	// 领域名称
	DomainName string `json:"domain_name" xml:"domain_name"`
}

func (o CreateProjectDomainRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProjectDomainRequestBody struct{}"
	}

	return strings.Join([]string{"CreateProjectDomainRequestBody", string(data)}, " ")
}
