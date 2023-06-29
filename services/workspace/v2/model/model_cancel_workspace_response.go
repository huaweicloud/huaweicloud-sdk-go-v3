package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelWorkspaceResponse Response Object
type CancelWorkspaceResponse struct {

	// 注销云办公服务的任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CancelWorkspaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelWorkspaceResponse struct{}"
	}

	return strings.Join([]string{"CancelWorkspaceResponse", string(data)}, " ")
}
