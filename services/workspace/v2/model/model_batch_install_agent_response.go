package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchInstallAgentResponse Response Object
type BatchInstallAgentResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchInstallAgentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchInstallAgentResponse struct{}"
	}

	return strings.Join([]string{"BatchInstallAgentResponse", string(data)}, " ")
}
