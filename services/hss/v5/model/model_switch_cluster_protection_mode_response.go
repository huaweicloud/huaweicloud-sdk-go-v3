package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchClusterProtectionModeResponse Response Object
type SwitchClusterProtectionModeResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SwitchClusterProtectionModeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchClusterProtectionModeResponse struct{}"
	}

	return strings.Join([]string{"SwitchClusterProtectionModeResponse", string(data)}, " ")
}
