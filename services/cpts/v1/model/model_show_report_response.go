package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowReportResponse struct {

	// code
	Code *string `json:"code,omitempty" xml:"code"`

	// message
	Message *string `json:"message,omitempty" xml:"message"`

	// extend
	Extend *string `json:"extend,omitempty" xml:"extend"`

	Result         *ReportInfo `json:"result,omitempty" xml:"result"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowReportResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReportResponse struct{}"
	}

	return strings.Join([]string{"ShowReportResponse", string(data)}, " ")
}
