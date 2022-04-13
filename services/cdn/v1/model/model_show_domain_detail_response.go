package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowDomainDetailResponse struct {
	Domain         *DomainsWithPort `json:"domain,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowDomainDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainDetailResponse", string(data)}, " ")
}
