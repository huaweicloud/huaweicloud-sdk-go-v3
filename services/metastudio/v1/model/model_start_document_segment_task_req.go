package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartDocumentSegmentTaskReq 开始分段任务请求信息。
type StartDocumentSegmentTaskReq struct {

	// 文档ID。
	DocumentId string `json:"document_id"`
}

func (o StartDocumentSegmentTaskReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartDocumentSegmentTaskReq struct{}"
	}

	return strings.Join([]string{"StartDocumentSegmentTaskReq", string(data)}, " ")
}
