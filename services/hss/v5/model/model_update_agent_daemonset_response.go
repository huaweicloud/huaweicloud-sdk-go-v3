package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAgentDaemonsetResponse Response Object
type UpdateAgentDaemonsetResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateAgentDaemonsetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAgentDaemonsetResponse struct{}"
	}

	return strings.Join([]string{"UpdateAgentDaemonsetResponse", string(data)}, " ")
}
