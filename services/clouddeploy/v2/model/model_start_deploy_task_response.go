package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type StartDeployTaskResponse struct {

	// 执行记录id
	Id *string `json:"id,omitempty" xml:"id"`

	// 部署任务id
	TaskId *string `json:"task_id,omitempty" xml:"task_id"`

	// 执行任务名称
	JobName *string `json:"job_name,omitempty" xml:"job_name"`

	// 部署任务和应用组件对应关系
	AppComponentList *[]AppComponentDao `json:"app_component_list,omitempty" xml:"app_component_list"`
	HttpStatusCode   int                `json:"-"`
}

func (o StartDeployTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartDeployTaskResponse struct{}"
	}

	return strings.Join([]string{"StartDeployTaskResponse", string(data)}, " ")
}
