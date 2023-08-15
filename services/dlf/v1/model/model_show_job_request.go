package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobRequest Request Object
type ShowJobRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// 作业名称.
	JobName string `json:"job_name"`
}

func (o ShowJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobRequest struct{}"
	}

	return strings.Join([]string{"ShowJobRequest", string(data)}, " ")
}
