package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDeployTaskByTemplateResponse Response Object
type CreateDeployTaskByTemplateResponse struct {

	// 应用名称
	TaskName *string `json:"task_name,omitempty"`

	// 部署任务id
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDeployTaskByTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDeployTaskByTemplateResponse struct{}"
	}

	return strings.Join([]string{"CreateDeployTaskByTemplateResponse", string(data)}, " ")
}
