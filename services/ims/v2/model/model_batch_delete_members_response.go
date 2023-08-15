package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteMembersResponse Response Object
type BatchDeleteMembersResponse struct {

	// 异步任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchDeleteMembersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteMembersResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteMembersResponse", string(data)}, " ")
}
