package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ExpandReplicationResponse struct {

	// 成功返回jobId信息
	JobId          *string `json:"job_id,omitempty" xml:"job_id"`
	HttpStatusCode int     `json:"-"`
}

func (o ExpandReplicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandReplicationResponse struct{}"
	}

	return strings.Join([]string{"ExpandReplicationResponse", string(data)}, " ")
}
