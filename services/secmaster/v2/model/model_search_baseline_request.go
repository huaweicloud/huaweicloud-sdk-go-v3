package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchBaselineRequest Request Object
type SearchBaselineRequest struct {

	// 工作空间id
	WorkspaceId string `json:"workspace_id"`

	// 语言，参考值：zh-CN、en-US
	XLanguage string `json:"X-Language"`

	// 内容类型
	ContentType string `json:"content-type"`

	Body *BaselineSearchRequestBody `json:"body,omitempty"`
}

func (o SearchBaselineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchBaselineRequest struct{}"
	}

	return strings.Join([]string{"SearchBaselineRequest", string(data)}, " ")
}
