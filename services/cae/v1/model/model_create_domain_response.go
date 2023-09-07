package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDomainResponse Response Object
type CreateDomainResponse struct {
	ApiVersion *ApiVersionObj `json:"api_version,omitempty"`

	// 域名列表。
	Items *[]DomainItem `json:"items,omitempty"`

	// API类型，固定值“Domain”，该值不可修改。
	Kind           *string `json:"kind,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDomainResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDomainResponse struct{}"
	}

	return strings.Join([]string{"CreateDomainResponse", string(data)}, " ")
}
