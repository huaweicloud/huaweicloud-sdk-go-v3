package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRiskInfoResponse Response Object
type ShowRiskInfoResponse struct {

	// 风险版本信息
	Risks *[]ShowRiskInfoEngineRiskDesc `json:"risks,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowRiskInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRiskInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowRiskInfoResponse", string(data)}, " ")
}
