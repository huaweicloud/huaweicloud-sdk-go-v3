package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PipelineBasicInfo struct {

	// devCloud项目id
	ProjectId string `json:"project_id"`

	// devCloud项目名字
	ProjectName string `json:"project_name"`

	// 流水线id
	PipelineId string `json:"pipeline_id"`

	// 流水线名字
	PipelineName string `json:"pipeline_name"`

	// 流水线创建人id
	CreatorId string `json:"creator_id"`

	// 流水线创建人名字
	CreatorName string `json:"creator_name"`

	// 流水线创建人id
	ExecutorId string `json:"executor_id"`

	// 流水线执行人名字
	ExecutorName string `json:"executor_name"`

	// 启动时间
	StartTime string `json:"start_time"`

	// 创建时间
	CreateTime string `json:"create_time"`

	// 用户是否关注流水线：true（关注），false（未关注）
	Watched string `json:"watched"`
}

func (o PipelineBasicInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineBasicInfo struct{}"
	}

	return strings.Join([]string{"PipelineBasicInfo", string(data)}, " ")
}
