package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetAccessPreviewRequest Request Object
type GetAccessPreviewRequest struct {

	// 分析器的唯一标识符。
	AnalyzerId string `json:"analyzer_id"`

	// 分析预览的唯一标识符。
	AccessPreviewId string `json:"access_preview_id"`
}

func (o GetAccessPreviewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetAccessPreviewRequest struct{}"
	}

	return strings.Join([]string{"GetAccessPreviewRequest", string(data)}, " ")
}
