package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunOnceRequest Request Object
type RunOnceRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// 作业名称.
	JobName string `json:"job_name"`

	Body *StartJobReq `json:"body,omitempty"`
}

func (o RunOnceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunOnceRequest struct{}"
	}

	return strings.Join([]string{"RunOnceRequest", string(data)}, " ")
}
