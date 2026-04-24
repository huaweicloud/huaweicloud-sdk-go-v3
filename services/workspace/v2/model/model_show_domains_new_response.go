package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDomainsNewResponse Response Object
type ShowDomainsNewResponse struct {

	// 统信域控列表。
	UosDomainList *[]UosDomainInfo `json:"uos_domain_list,omitempty"`

	// 域信息。
	DomainInfos    *[]AdDomain `json:"domain_infos,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowDomainsNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainsNewResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainsNewResponse", string(data)}, " ")
}
