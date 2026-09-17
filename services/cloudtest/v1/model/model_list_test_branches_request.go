package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTestBranchesRequest Request Object
type ListTestBranchesRequest struct {

	// 项目ID（云龙场景，传入微服务ID）
	ProjectUuid string `json:"project_uuid"`

	// 排序字段
	SortField *string `json:"sort_field,omitempty"`

	// 排序方式
	SortType *string `json:"sort_type,omitempty"`
}

func (o ListTestBranchesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTestBranchesRequest struct{}"
	}

	return strings.Join([]string{"ListTestBranchesRequest", string(data)}, " ")
}
