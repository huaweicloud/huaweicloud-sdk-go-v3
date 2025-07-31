package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteAgentDaemonsetResponse Response Object
type BatchDeleteAgentDaemonsetResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteAgentDaemonsetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteAgentDaemonsetResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteAgentDaemonsetResponse", string(data)}, " ")
}
