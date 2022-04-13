package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchAddMembersResponse struct {
	// 异步任务ID。

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchAddMembersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddMembersResponse struct{}"
	}

	return strings.Join([]string{"BatchAddMembersResponse", string(data)}, " ")
}
