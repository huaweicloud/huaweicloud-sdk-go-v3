package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListParseDomainDetailsRequest struct {

	// 域名
	DomainName string `json:"domain_name"`
}

func (o ListParseDomainDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListParseDomainDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListParseDomainDetailsRequest", string(data)}, " ")
}
