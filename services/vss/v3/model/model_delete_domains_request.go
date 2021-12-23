package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteDomainsRequest struct {
	// 域名

	DomainName string `json:"domain_name"`
}

func (o DeleteDomainsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDomainsRequest struct{}"
	}

	return strings.Join([]string{"DeleteDomainsRequest", string(data)}, " ")
}
