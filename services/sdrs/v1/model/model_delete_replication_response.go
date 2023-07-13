package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteReplicationResponse Response Object
type DeleteReplicationResponse struct {

	// 成功返回jobId信息
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteReplicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteReplicationResponse struct{}"
	}

	return strings.Join([]string{"DeleteReplicationResponse", string(data)}, " ")
}
