package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHealthReportTemplateResponse Response Object
type ListHealthReportTemplateResponse struct {

	// 健康报告模板列表（系统预设 + 自定义）。
	Templates      *[]HealthReportTemplateDto `json:"templates,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListHealthReportTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHealthReportTemplateResponse struct{}"
	}

	return strings.Join([]string{"ListHealthReportTemplateResponse", string(data)}, " ")
}
