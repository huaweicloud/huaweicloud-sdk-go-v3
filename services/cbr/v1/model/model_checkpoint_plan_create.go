package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CheckpointPlanCreate struct {

	// 存储库id
	Id string `json:"id" xml:"id"`

	// 存储库名称
	Name string `json:"name" xml:"name"`

	// 备份对象
	Resources *[]CheckpointResourceResp `json:"resources,omitempty" xml:"resources"`

	// 备份时跳过的资源列表
	SkippedResources *[]CheckpointCreateSkippedResource `json:"skipped_resources,omitempty" xml:"skipped_resources"`
}

func (o CheckpointPlanCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckpointPlanCreate struct{}"
	}

	return strings.Join([]string{"CheckpointPlanCreate", string(data)}, " ")
}
