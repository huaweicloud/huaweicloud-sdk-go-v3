package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetachProtectedInstanceReplicationResponse Response Object
type DetachProtectedInstanceReplicationResponse struct {

	// 成功返回jobId信息
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DetachProtectedInstanceReplicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachProtectedInstanceReplicationResponse struct{}"
	}

	return strings.Join([]string{"DetachProtectedInstanceReplicationResponse", string(data)}, " ")
}
