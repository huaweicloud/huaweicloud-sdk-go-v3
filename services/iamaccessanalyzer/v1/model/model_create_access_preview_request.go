package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAccessPreviewRequest Request Object
type CreateAccessPreviewRequest struct {

	// 分析器的唯一标识符。
	AnalyzerId string `json:"analyzer_id"`

	Body *CreateAccessPreviewReqBody `json:"body,omitempty"`
}

func (o CreateAccessPreviewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAccessPreviewRequest struct{}"
	}

	return strings.Join([]string{"CreateAccessPreviewRequest", string(data)}, " ")
}
