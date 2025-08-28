package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDocumentSegmentReq 修改文档分段信息请求。
type UpdateDocumentSegmentReq struct {

	// 文档ID。
	DocumentId string `json:"document_id"`

	// 文档分段文本ID。
	Id string `json:"id"`

	// 分段序号
	TextIndex *int32 `json:"text_index,omitempty"`

	// 文档分段文本。
	Text *string `json:"text,omitempty"`
}

func (o UpdateDocumentSegmentReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDocumentSegmentReq struct{}"
	}

	return strings.Join([]string{"UpdateDocumentSegmentReq", string(data)}, " ")
}
