package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSystemDefaultDomainConfigResponse Response Object
type UpdateSystemDefaultDomainConfigResponse struct {
	Loadbalancer   *DnsConfigResponseBody `json:"loadbalancer,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o UpdateSystemDefaultDomainConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSystemDefaultDomainConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateSystemDefaultDomainConfigResponse", string(data)}, " ")
}
