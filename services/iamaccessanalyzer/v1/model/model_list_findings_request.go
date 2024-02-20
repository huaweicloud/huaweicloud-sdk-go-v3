package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFindingsRequest Request Object
type ListFindingsRequest struct {

	// 分析器的唯一标识符。
	AnalyzerId string `json:"analyzer_id"`

	Body *ListFindingsReqBody `json:"body,omitempty"`
}

func (o ListFindingsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFindingsRequest struct{}"
	}

	return strings.Join([]string{"ListFindingsRequest", string(data)}, " ")
}
