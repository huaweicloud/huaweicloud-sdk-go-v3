package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTenantConfigRequest Request Object
type UpdateTenantConfigRequest struct {
	Body *UpdateTenantConfigReq `json:"body,omitempty"`
}

func (o UpdateTenantConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTenantConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateTenantConfigRequest", string(data)}, " ")
}
