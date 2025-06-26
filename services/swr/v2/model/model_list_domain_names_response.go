package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDomainNamesResponse Response Object
type ListDomainNamesResponse struct {

	// 域名信息
	DomainNameInfos *[]DomainNameInfo `json:"domain_name_infos,omitempty"`
	HttpStatusCode  int               `json:"-"`
}

func (o ListDomainNamesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainNamesResponse struct{}"
	}

	return strings.Join([]string{"ListDomainNamesResponse", string(data)}, " ")
}
