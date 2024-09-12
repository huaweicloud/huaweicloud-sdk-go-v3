package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDomainsRequest Request Object
type UpdateDomainsRequest struct {

	// 域名
	DomainId string `json:"domain_id"`

	Body *UpdateDomainRequestBody `json:"body,omitempty"`
}

func (o UpdateDomainsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainsRequest struct{}"
	}

	return strings.Join([]string{"UpdateDomainsRequest", string(data)}, " ")
}
