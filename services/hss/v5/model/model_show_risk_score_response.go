package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRiskScoreResponse Response Object
type ShowRiskScoreResponse struct {

	// 安全评分
	Score *int32 `json:"score,omitempty"`

	// 安全风险个数
	RiskNum *int32 `json:"risk_num,omitempty"`

	// 安全风险主机个数
	RiskServerNum  *int32 `json:"risk_server_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowRiskScoreResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRiskScoreResponse struct{}"
	}

	return strings.Join([]string{"ShowRiskScoreResponse", string(data)}, " ")
}
