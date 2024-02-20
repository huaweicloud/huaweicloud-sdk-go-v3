package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateArchiveRuleRequest Request Object
type CreateArchiveRuleRequest struct {

	// 分析器的唯一标识符。
	AnalyzerId string `json:"analyzer_id"`

	Body *CreateArchiveRuleReqBody `json:"body,omitempty"`
}

func (o CreateArchiveRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateArchiveRuleRequest struct{}"
	}

	return strings.Join([]string{"CreateArchiveRuleRequest", string(data)}, " ")
}
