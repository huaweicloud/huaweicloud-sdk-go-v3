package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRiskInfoResponse Response Object
type ShowRiskInfoResponse struct {

	// 内核风险提醒信息。
	Risks          *[]Risks `json:"risks,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ShowRiskInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRiskInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowRiskInfoResponse", string(data)}, " ")
}
