package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDomainNameResponse Response Object
type UpdateDomainNameResponse struct {
	DomainNameInfo *DomainNameInfo `json:"domain_name_info,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o UpdateDomainNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainNameResponse struct{}"
	}

	return strings.Join([]string{"UpdateDomainNameResponse", string(data)}, " ")
}
