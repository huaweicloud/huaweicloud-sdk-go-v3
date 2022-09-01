package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type JobEntitiesResult struct {

	// 镜像ID。
	ImageId *string `json:"image_id,omitempty" xml:"image_id"`

	// 项目ID。
	ProjectId *string `json:"project_id,omitempty" xml:"project_id"`

	// 任务状态。
	Status *string `json:"status,omitempty" xml:"status"`
}

func (o JobEntitiesResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobEntitiesResult struct{}"
	}

	return strings.Join([]string{"JobEntitiesResult", string(data)}, " ")
}
