package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRelationTableResponse Response Object
type ListRelationTableResponse struct {

	// 关系信息总数
	Total *int32 `json:"total,omitempty"`

	// 当前页
	CurrentPage *int32 `json:"current_page,omitempty"`

	// 关系信息列表
	TableList      *[]RelationSimpleInfo `json:"table_list,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListRelationTableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRelationTableResponse struct{}"
	}

	return strings.Join([]string{"ListRelationTableResponse", string(data)}, " ")
}
