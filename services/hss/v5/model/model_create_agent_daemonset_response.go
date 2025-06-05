package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAgentDaemonsetResponse Response Object
type CreateAgentDaemonsetResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateAgentDaemonsetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAgentDaemonsetResponse struct{}"
	}

	return strings.Join([]string{"CreateAgentDaemonsetResponse", string(data)}, " ")
}
