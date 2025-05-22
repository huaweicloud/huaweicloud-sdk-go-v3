package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowReportSummaryResponse Response Object
type ShowReportSummaryResponse struct {

	// 状态
	Status *string `json:"status,omitempty"`

	// 错误
	Error *interface{} `json:"error,omitempty"`

	Result         *ShowReportSummaryBodyResult `json:"result,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ShowReportSummaryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReportSummaryResponse struct{}"
	}

	return strings.Join([]string{"ShowReportSummaryResponse", string(data)}, " ")
}
