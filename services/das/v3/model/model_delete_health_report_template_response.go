package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHealthReportTemplateResponse Response Object
type DeleteHealthReportTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteHealthReportTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHealthReportTemplateResponse struct{}"
	}

	return strings.Join([]string{"DeleteHealthReportTemplateResponse", string(data)}, " ")
}
