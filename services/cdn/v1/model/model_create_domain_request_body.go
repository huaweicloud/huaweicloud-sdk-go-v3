package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDomainRequestBody 域名对象
type CreateDomainRequestBody struct {
	Domain *DomainBody `json:"domain"`
}

func (o CreateDomainRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDomainRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDomainRequestBody", string(data)}, " ")
}
