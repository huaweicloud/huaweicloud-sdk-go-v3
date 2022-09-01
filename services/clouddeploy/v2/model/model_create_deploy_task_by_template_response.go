package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateDeployTaskByTemplateResponse struct {

	// 部署任务名称
	TaskName *string `json:"task_name,omitempty" xml:"task_name"`

	// 部署任务id
	TaskId         *string `json:"task_id,omitempty" xml:"task_id"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDeployTaskByTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDeployTaskByTemplateResponse struct{}"
	}

	return strings.Join([]string{"CreateDeployTaskByTemplateResponse", string(data)}, " ")
}
