package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDomainRequest Request Object
type UpdateDomainRequest struct {

	// 域名ID
	DomainId string `json:"domain_id"`

	Body *DomainRealServerInfo `json:"body,omitempty"`
}

func (o UpdateDomainRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainRequest struct{}"
	}

	return strings.Join([]string{"UpdateDomainRequest", string(data)}, " ")
}
