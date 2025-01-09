package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetMaintenanceModeResponse Response Object
type SetMaintenanceModeResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SetMaintenanceModeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetMaintenanceModeResponse struct{}"
	}

	return strings.Join([]string{"SetMaintenanceModeResponse", string(data)}, " ")
}
