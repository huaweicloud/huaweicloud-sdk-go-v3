package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartDeployTaskResponse Response Object
type StartDeployTaskResponse struct {

	// 部署记录id
	Id *string `json:"id,omitempty"`

	// 部署任务id
	TaskId *string `json:"task_id,omitempty"`

	// 执行任务名称
	JobName *string `json:"job_name,omitempty"`

	// 应用和AOM应用组件对应关系
	AppComponentList *[]AppComponentDao `json:"app_component_list,omitempty"`
	HttpStatusCode   int                `json:"-"`
}

func (o StartDeployTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartDeployTaskResponse struct{}"
	}

	return strings.Join([]string{"StartDeployTaskResponse", string(data)}, " ")
}
