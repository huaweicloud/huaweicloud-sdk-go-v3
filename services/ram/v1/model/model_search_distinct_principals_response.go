package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type SearchDistinctPrincipalsResponse struct {

	// 不同角色的信息列表。
	DistinctSharedPrincipals *[]DistinctSharedPrincipal `json:"distinct_shared_principals,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o SearchDistinctPrincipalsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchDistinctPrincipalsResponse struct{}"
	}

	return strings.Join([]string{"SearchDistinctPrincipalsResponse", string(data)}, " ")
}
