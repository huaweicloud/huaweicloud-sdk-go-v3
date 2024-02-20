package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPreviewFindingsRequest Request Object
type ListPreviewFindingsRequest struct {

	// 分析器的唯一标识符。
	AnalyzerId string `json:"analyzer_id"`

	// 分析预览的唯一标识符。
	AccessPreviewId string `json:"access_preview_id"`

	Body *ListPreviewFindingsReqBody `json:"body,omitempty"`
}

func (o ListPreviewFindingsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPreviewFindingsRequest struct{}"
	}

	return strings.Join([]string{"ListPreviewFindingsRequest", string(data)}, " ")
}
