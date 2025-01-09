package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetMaintenanceModeRequest Request Object
type SetMaintenanceModeRequest struct {
	Body *BatchSetMaintenanceModeReq `json:"body,omitempty"`
}

func (o SetMaintenanceModeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetMaintenanceModeRequest struct{}"
	}

	return strings.Join([]string{"SetMaintenanceModeRequest", string(data)}, " ")
}
