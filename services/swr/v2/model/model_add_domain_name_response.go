package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddDomainNameResponse Response Object
type AddDomainNameResponse struct {
	DomainNameInfo *DomainNameInfo `json:"domain_name_info,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o AddDomainNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDomainNameResponse struct{}"
	}

	return strings.Join([]string{"AddDomainNameResponse", string(data)}, " ")
}
