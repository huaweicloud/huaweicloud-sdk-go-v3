package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAllBranchesRequest Request Object
type ListAllBranchesRequest struct {

	// 项目ID（云龙场景，传入微服务ID）
	ProjectUuid string `json:"project_uuid"`

	// 排序字段
	SortField *string `json:"sort_field,omitempty"`

	// 排序方式
	SortType *string `json:"sort_type,omitempty"`
}

func (o ListAllBranchesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllBranchesRequest struct{}"
	}

	return strings.Join([]string{"ListAllBranchesRequest", string(data)}, " ")
}
