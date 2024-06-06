package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAccessPreviewRequest Request Object
type ShowAccessPreviewRequest struct {

	// 分析器的唯一标识符。
	AnalyzerId string `json:"analyzer_id"`

	// 访问预览的唯一标识符。
	AccessPreviewId string `json:"access_preview_id"`
}

func (o ShowAccessPreviewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAccessPreviewRequest struct{}"
	}

	return strings.Join([]string{"ShowAccessPreviewRequest", string(data)}, " ")
}
