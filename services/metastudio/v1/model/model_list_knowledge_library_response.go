package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListKnowledgeLibraryResponse Response Object
type ListKnowledgeLibraryResponse struct {

	// 与第一条数据的偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 页面大小
	Limit *int32 `json:"limit,omitempty"`

	// 总数量
	Count *int32 `json:"count,omitempty"`

	// 知识库信息
	Data *[]KnowledgeLibraryInfo `json:"data,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListKnowledgeLibraryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKnowledgeLibraryResponse struct{}"
	}

	return strings.Join([]string{"ListKnowledgeLibraryResponse", string(data)}, " ")
}
