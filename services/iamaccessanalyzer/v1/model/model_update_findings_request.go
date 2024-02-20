package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFindingsRequest Request Object
type UpdateFindingsRequest struct {

	// 分析器的唯一标识符。
	AnalyzerId string `json:"analyzer_id"`

	Body *UpdateFindingsReqBody `json:"body,omitempty"`
}

func (o UpdateFindingsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFindingsRequest struct{}"
	}

	return strings.Join([]string{"UpdateFindingsRequest", string(data)}, " ")
}
