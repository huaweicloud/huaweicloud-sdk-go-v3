package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowInstanceStatusResponse struct {

	// 实例ID
	TaskId *string `json:"task_id,omitempty" xml:"task_id"`

	// 实例创建状态，取值和含义:initializing：初始化中；processing：处理中；finished：已完成；succeeded：成功
	TaskStatus *string `json:"task_status,omitempty" xml:"task_status"`

	// 流水线ID
	PipelineId *string `json:"pipeline_id,omitempty" xml:"pipeline_id"`

	// 流水线名字
	PipelineName *string `json:"pipeline_name,omitempty" xml:"pipeline_name"`

	// 流水线详情页面url
	PipelineUrl    *string `json:"pipeline_url,omitempty" xml:"pipeline_url"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowInstanceStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceStatusResponse", string(data)}, " ")
}
