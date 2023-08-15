package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowReportResponse Response Object
type ShowReportResponse struct {

	// code
	Code *string `json:"code,omitempty"`

	// message
	Message *string `json:"message,omitempty"`

	// extend
	Extend *string `json:"extend,omitempty"`

	Result         *ReportInfo `json:"result,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowReportResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReportResponse struct{}"
	}

	return strings.Join([]string{"ShowReportResponse", string(data)}, " ")
}
