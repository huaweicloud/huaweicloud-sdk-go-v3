package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowTablePreviewRequest struct {

	// 表ID。
	TableId string `json:"table_id" xml:"table_id"`
}

func (o ShowTablePreviewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTablePreviewRequest struct{}"
	}

	return strings.Join([]string{"ShowTablePreviewRequest", string(data)}, " ")
}
