package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowDomainDetailByNameResponse struct {
	Domain         *DomainsWithPort `json:"domain,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowDomainDetailByNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainDetailByNameResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainDetailByNameResponse", string(data)}, " ")
}
