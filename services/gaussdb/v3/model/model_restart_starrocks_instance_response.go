package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestartStarrocksInstanceResponse Response Object
type RestartStarrocksInstanceResponse struct {

	// 工作流ID。
	WorkflowId     *string `json:"workflow_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RestartStarrocksInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartStarrocksInstanceResponse struct{}"
	}

	return strings.Join([]string{"RestartStarrocksInstanceResponse", string(data)}, " ")
}
