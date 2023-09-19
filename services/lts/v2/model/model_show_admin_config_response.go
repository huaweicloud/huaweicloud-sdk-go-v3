package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAdminConfigResponse Response Object
type ShowAdminConfigResponse struct {
	LogConvergeSwitch *bool `json:"log_converge_switch,omitempty"`
	HttpStatusCode    int   `json:"-"`
}

func (o ShowAdminConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAdminConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowAdminConfigResponse", string(data)}, " ")
}
