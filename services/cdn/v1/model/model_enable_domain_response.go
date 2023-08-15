package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableDomainResponse Response Object
type EnableDomainResponse struct {
	Domain         *DomainsWithPort `json:"domain,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o EnableDomainResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableDomainResponse struct{}"
	}

	return strings.Join([]string{"EnableDomainResponse", string(data)}, " ")
}
