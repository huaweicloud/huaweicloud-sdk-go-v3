package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateUserDefinedDomainConfigResponse Response Object
type UpdateUserDefinedDomainConfigResponse struct {
	Loadbalancer   *DnsConfigResponseBody `json:"loadbalancer,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o UpdateUserDefinedDomainConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserDefinedDomainConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateUserDefinedDomainConfigResponse", string(data)}, " ")
}
