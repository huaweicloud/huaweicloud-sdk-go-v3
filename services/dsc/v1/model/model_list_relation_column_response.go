package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRelationColumnResponse struct {

	// 关系信息总数
	Total *int32 `json:"total,omitempty"`

	// 关系信息列表
	ColumnList     *[]RelationSimpleInfo `json:"column_list,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListRelationColumnResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRelationColumnResponse struct{}"
	}

	return strings.Join([]string{"ListRelationColumnResponse", string(data)}, " ")
}
