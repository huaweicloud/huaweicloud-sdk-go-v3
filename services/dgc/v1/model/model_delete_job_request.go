package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteJobRequest Request Object
type DeleteJobRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// 作业名称.
	JobName string `json:"job_name"`

	Body *DeleteReq `json:"body,omitempty"`
}

func (o DeleteJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteJobRequest struct{}"
	}

	return strings.Join([]string{"DeleteJobRequest", string(data)}, " ")
}
