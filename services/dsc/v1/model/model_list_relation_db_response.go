package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRelationDbResponse struct {

	// 关系信息总数
	Total *int32 `json:"total,omitempty"`

	// 关系信息列表
	DbList         *[]RelationSimpleInfo `json:"db_list,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListRelationDbResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRelationDbResponse struct{}"
	}

	return strings.Join([]string{"ListRelationDbResponse", string(data)}, " ")
}
