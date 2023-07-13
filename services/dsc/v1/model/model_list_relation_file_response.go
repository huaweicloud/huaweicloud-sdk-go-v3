package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRelationFileResponse Response Object
type ListRelationFileResponse struct {

	// 关系信息总数
	Total *int32 `json:"total,omitempty"`

	// 当前页
	CurrentPage *int32 `json:"current_page,omitempty"`

	// 关系信息列表
	FileList       *[]RelationSimpleInfo `json:"file_list,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListRelationFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRelationFileResponse struct{}"
	}

	return strings.Join([]string{"ListRelationFileResponse", string(data)}, " ")
}
