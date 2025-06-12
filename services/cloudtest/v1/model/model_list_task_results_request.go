package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTaskResultsRequest Request Object
type ListTaskResultsRequest struct {

	// 项目id
	ProjectUuid string `json:"project_uuid"`

	// 任务URI
	TaskUri string `json:"task_uri"`

	// 测试计划id
	IteratorUri *string `json:"iterator_uri,omitempty"`

	// 当前页数
	PageNo string `json:"page_no"`

	// 每页多少记录
	PageSize string `json:"page_size"`

	// 发布版本
	ReleaseDev *string `json:"release_dev,omitempty"`
}

func (o ListTaskResultsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTaskResultsRequest struct{}"
	}

	return strings.Join([]string{"ListTaskResultsRequest", string(data)}, " ")
}
