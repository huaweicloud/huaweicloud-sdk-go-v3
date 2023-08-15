package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateReplicationResponse Response Object
type CreateReplicationResponse struct {

	// 成功返回jobId信息
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateReplicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateReplicationResponse struct{}"
	}

	return strings.Join([]string{"CreateReplicationResponse", string(data)}, " ")
}
