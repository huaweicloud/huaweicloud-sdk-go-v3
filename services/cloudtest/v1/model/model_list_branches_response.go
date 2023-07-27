package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBranchesResponse Response Object
type ListBranchesResponse struct {

	// 起始记录数 大于 实际总条数时， 值为0， 分页请求才有此值
	Total *int32 `json:"total,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 实际的数据类型：单个对象，集合 或 NULL
	Values         *[]ExternalBranchInfoVo `json:"values,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListBranchesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBranchesResponse struct{}"
	}

	return strings.Join([]string{"ListBranchesResponse", string(data)}, " ")
}
