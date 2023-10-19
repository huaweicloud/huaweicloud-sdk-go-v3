package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDomainSetsResponse Response Object
type ListDomainSetsResponse struct {
	Data           *ListDomainsetsResponseData `json:"data,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListDomainSetsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainSetsResponse struct{}"
	}

	return strings.Join([]string{"ListDomainSetsResponse", string(data)}, " ")
}
