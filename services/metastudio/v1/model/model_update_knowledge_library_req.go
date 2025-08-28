package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateKnowledgeLibraryReq 修改知识库请求。
type UpdateKnowledgeLibraryReq struct {

	// 知识库名称。
	Name *string `json:"name,omitempty"`

	// 知识库召回数
	Topk *int32 `json:"topk,omitempty"`

	// 知识库召回得分
	Score *float64 `json:"score,omitempty"`
}

func (o UpdateKnowledgeLibraryReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKnowledgeLibraryReq struct{}"
	}

	return strings.Join([]string{"UpdateKnowledgeLibraryReq", string(data)}, " ")
}
