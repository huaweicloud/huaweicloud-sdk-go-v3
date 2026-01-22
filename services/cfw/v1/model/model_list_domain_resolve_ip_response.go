package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDomainResolveIpResponse Response Object
type ListDomainResolveIpResponse struct {
	Data           *ParseIpListResponse `json:"data,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListDomainResolveIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainResolveIpResponse struct{}"
	}

	return strings.Join([]string{"ListDomainResolveIpResponse", string(data)}, " ")
}
