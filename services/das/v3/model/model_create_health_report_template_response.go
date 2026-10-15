package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateHealthReportTemplateResponse Response Object
type CreateHealthReportTemplateResponse struct {

	// 新创建的模板ID。
	TemplateId     *int64 `json:"template_id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateHealthReportTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHealthReportTemplateResponse struct{}"
	}

	return strings.Join([]string{"CreateHealthReportTemplateResponse", string(data)}, " ")
}
