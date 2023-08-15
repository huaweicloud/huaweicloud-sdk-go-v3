package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchSharedPrincipalsResponse Response Object
type SearchSharedPrincipalsResponse struct {

	// 资源使用者的详细信息列表。
	SharedPrincipals *[]SharedPrincipal `json:"shared_principals,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o SearchSharedPrincipalsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchSharedPrincipalsResponse struct{}"
	}

	return strings.Join([]string{"SearchSharedPrincipalsResponse", string(data)}, " ")
}
