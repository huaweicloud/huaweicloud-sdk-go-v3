package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAgentDaemonsetResponse Response Object
type DeleteAgentDaemonsetResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAgentDaemonsetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAgentDaemonsetResponse struct{}"
	}

	return strings.Join([]string{"DeleteAgentDaemonsetResponse", string(data)}, " ")
}
