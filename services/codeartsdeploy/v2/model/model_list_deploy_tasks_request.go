package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDeployTasksRequest Request Object
type ListDeployTasksRequest struct {

	// 项目id
	ProjectId string `json:"project_id"`

	// 分页页码， 表示从此页开始查询， page大于等于1
	Page int32 `json:"page"`

	// 每页显示的条目数量，size小于等于100
	Size int32 `json:"size"`
}

func (o ListDeployTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDeployTasksRequest struct{}"
	}

	return strings.Join([]string{"ListDeployTasksRequest", string(data)}, " ")
}
