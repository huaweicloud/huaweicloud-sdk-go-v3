package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAccessPreviewFindingsResponse Response Object
type ListAccessPreviewFindingsResponse struct {

	// 访问预览生成的分析结果列表。
	Findings *[]PreviewFinding `json:"findings,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListAccessPreviewFindingsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccessPreviewFindingsResponse struct{}"
	}

	return strings.Join([]string{"ListAccessPreviewFindingsResponse", string(data)}, " ")
}
