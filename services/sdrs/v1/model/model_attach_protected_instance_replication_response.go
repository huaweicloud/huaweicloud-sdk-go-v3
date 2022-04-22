package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type AttachProtectedInstanceReplicationResponse struct {

	// 成功返回jobId信息
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AttachProtectedInstanceReplicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachProtectedInstanceReplicationResponse struct{}"
	}

	return strings.Join([]string{"AttachProtectedInstanceReplicationResponse", string(data)}, " ")
}
