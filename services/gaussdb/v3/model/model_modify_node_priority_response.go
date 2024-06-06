package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyNodePriorityResponse Response Object
type ModifyNodePriorityResponse struct {

	// 工作流ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ModifyNodePriorityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyNodePriorityResponse struct{}"
	}

	return strings.Join([]string{"ModifyNodePriorityResponse", string(data)}, " ")
}
