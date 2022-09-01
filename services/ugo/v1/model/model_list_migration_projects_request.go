package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListMigrationProjectsRequest struct {

	// 分页查询的偏移量。
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`
}

func (o ListMigrationProjectsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMigrationProjectsRequest struct{}"
	}

	return strings.Join([]string{"ListMigrationProjectsRequest", string(data)}, " ")
}
