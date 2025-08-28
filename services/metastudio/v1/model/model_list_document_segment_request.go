package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDocumentSegmentRequest Request Object
type ListDocumentSegmentRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 使用AK/SK方式认证时必选，携带项目ID信息。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 第三方用户ID。不允许输入中文。
	XAppUserId *string `json:"X-App-UserId,omitempty"`

	// 偏移量，表示从此偏移量开始查询。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 文档ID。
	DocumentId string `json:"document_id"`
}

func (o ListDocumentSegmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDocumentSegmentRequest struct{}"
	}

	return strings.Join([]string{"ListDocumentSegmentRequest", string(data)}, " ")
}
