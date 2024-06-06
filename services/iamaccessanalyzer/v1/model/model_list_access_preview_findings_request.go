package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAccessPreviewFindingsRequest Request Object
type ListAccessPreviewFindingsRequest struct {

	// 分析器的唯一标识符。
	AnalyzerId string `json:"analyzer_id"`

	// 访问预览的唯一标识符。
	AccessPreviewId string `json:"access_preview_id"`

	Body *ListPreviewFindingsReqBody `json:"body,omitempty"`
}

func (o ListAccessPreviewFindingsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccessPreviewFindingsRequest struct{}"
	}

	return strings.Join([]string{"ListAccessPreviewFindingsRequest", string(data)}, " ")
}
