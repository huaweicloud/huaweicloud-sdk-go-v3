package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryCaseResultInfo struct {

	// 结果URI
	Uri *string `json:"uri,omitempty"`

	// 测试任务URI
	TaskUri *string `json:"task_uri,omitempty"`

	// 执行id
	TaskId *string `json:"task_id,omitempty"`

	// 版本号
	ReleaseDev *string `json:"release_dev,omitempty"`

	// 当前页数
	PageNo *int32 `json:"page_no,omitempty"`

	// 每页条数
	PageSize *int32 `json:"page_size,omitempty"`

	// 排序字段
	SortField *string `json:"sort_field,omitempty"`

	// 排序类型
	SortType *string `json:"sort_type,omitempty"`
}

func (o QueryCaseResultInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryCaseResultInfo struct{}"
	}

	return strings.Join([]string{"QueryCaseResultInfo", string(data)}, " ")
}
