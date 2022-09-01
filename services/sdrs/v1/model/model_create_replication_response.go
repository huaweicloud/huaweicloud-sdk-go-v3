package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateReplicationResponse struct {

	// 成功返回jobId信息
	JobId          *string `json:"job_id,omitempty" xml:"job_id"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateReplicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateReplicationResponse struct{}"
	}

	return strings.Join([]string{"CreateReplicationResponse", string(data)}, " ")
}
