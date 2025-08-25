package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProjectSettingsInheritCfgResponse Response Object
type ShowProjectSettingsInheritCfgResponse struct {
	Body           *[]ProjectSettingsInheritCfgDto `json:"body,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ShowProjectSettingsInheritCfgResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectSettingsInheritCfgResponse struct{}"
	}

	return strings.Join([]string{"ShowProjectSettingsInheritCfgResponse", string(data)}, " ")
}
