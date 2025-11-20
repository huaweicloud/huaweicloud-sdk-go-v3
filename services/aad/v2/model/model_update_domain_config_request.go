package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDomainConfigRequest Request Object
type UpdateDomainConfigRequest struct {

	// 域名id
	DomainId string `json:"domain_id"`

	Body *UpdateDomainConfigRequestBody `json:"body,omitempty"`
}

func (o UpdateDomainConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateDomainConfigRequest", string(data)}, " ")
}
