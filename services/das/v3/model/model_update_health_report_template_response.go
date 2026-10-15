package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHealthReportTemplateResponse Response Object
type UpdateHealthReportTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateHealthReportTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHealthReportTemplateResponse struct{}"
	}

	return strings.Join([]string{"UpdateHealthReportTemplateResponse", string(data)}, " ")
}
