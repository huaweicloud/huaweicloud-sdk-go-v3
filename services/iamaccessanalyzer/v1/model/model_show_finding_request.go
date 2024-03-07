package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFindingRequest Request Object
type ShowFindingRequest struct {

	// 分析器的唯一标识符。
	AnalyzerId string `json:"analyzer_id"`

	// 访问分析结果的唯一标识符。
	FindingId string `json:"finding_id"`
}

func (o ShowFindingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFindingRequest struct{}"
	}

	return strings.Join([]string{"ShowFindingRequest", string(data)}, " ")
}
