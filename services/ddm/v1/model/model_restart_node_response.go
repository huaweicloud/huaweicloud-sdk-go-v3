package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestartNodeResponse Response Object
type RestartNodeResponse struct {

	// 任务id
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RestartNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartNodeResponse struct{}"
	}

	return strings.Join([]string{"RestartNodeResponse", string(data)}, " ")
}
