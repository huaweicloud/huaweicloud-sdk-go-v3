package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchDbAgentJobResponse Response Object
type SwitchDbAgentJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SwitchDbAgentJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchDbAgentJobResponse struct{}"
	}

	return strings.Join([]string{"SwitchDbAgentJobResponse", string(data)}, " ")
}
