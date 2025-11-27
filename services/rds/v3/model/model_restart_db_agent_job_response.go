package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestartDbAgentJobResponse Response Object
type RestartDbAgentJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RestartDbAgentJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartDbAgentJobResponse struct{}"
	}

	return strings.Join([]string{"RestartDbAgentJobResponse", string(data)}, " ")
}
