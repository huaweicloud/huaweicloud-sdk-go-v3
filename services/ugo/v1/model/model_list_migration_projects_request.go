package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMigrationProjectsRequest Request Object
type ListMigrationProjectsRequest struct {

	// 分页查询的偏移量。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListMigrationProjectsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMigrationProjectsRequest struct{}"
	}

	return strings.Join([]string{"ListMigrationProjectsRequest", string(data)}, " ")
}
