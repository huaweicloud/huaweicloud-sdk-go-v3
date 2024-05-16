package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestartStarrocksNodeResponse Response Object
type RestartStarrocksNodeResponse struct {

	// 工作流ID。
	WorkflowId     *string `json:"workflow_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RestartStarrocksNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartStarrocksNodeResponse struct{}"
	}

	return strings.Join([]string{"RestartStarrocksNodeResponse", string(data)}, " ")
}
