package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetResourcesRenewConfigRequest Request Object
type SetResourcesRenewConfigRequest struct {
	Body *RenewResourceConfigReq `json:"body,omitempty"`
}

func (o SetResourcesRenewConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetResourcesRenewConfigRequest struct{}"
	}

	return strings.Join([]string{"SetResourcesRenewConfigRequest", string(data)}, " ")
}
