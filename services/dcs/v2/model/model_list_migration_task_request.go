package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListMigrationTaskRequest struct {

	// 偏移量，表示从此偏移量开始查询， offset大于等于0。
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 迁移任务名称。
	Name *string `json:"name,omitempty" xml:"name"`
}

func (o ListMigrationTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMigrationTaskRequest struct{}"
	}

	return strings.Join([]string{"ListMigrationTaskRequest", string(data)}, " ")
}
