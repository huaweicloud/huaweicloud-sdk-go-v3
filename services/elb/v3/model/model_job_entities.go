package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobEntities 子任务关联的资源列表
type JobEntities struct {

	// 子任务关联的资源ID
	ResourceId *string `json:"resource_id,omitempty"`

	// 子任务关联的资源类型
	ResourceType *string `json:"resource_type,omitempty"`
}

func (o JobEntities) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobEntities struct{}"
	}

	return strings.Join([]string{"JobEntities", string(data)}, " ")
}
