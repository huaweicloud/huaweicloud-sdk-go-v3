package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDomainsRequest Request Object
type DeleteDomainsRequest struct {

	// 网站域名
	DomainName string `json:"domain_name"`
}

func (o DeleteDomainsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDomainsRequest struct{}"
	}

	return strings.Join([]string{"DeleteDomainsRequest", string(data)}, " ")
}
