package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstallAgentResponse Response Object
type InstallAgentResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o InstallAgentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstallAgentResponse struct{}"
	}

	return strings.Join([]string{"InstallAgentResponse", string(data)}, " ")
}
