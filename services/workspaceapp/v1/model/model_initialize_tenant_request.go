package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InitializeTenantRequest Request Object
type InitializeTenantRequest struct {
	Body *InitializeTenantReq `json:"body,omitempty"`
}

func (o InitializeTenantRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InitializeTenantRequest struct{}"
	}

	return strings.Join([]string{"InitializeTenantRequest", string(data)}, " ")
}
