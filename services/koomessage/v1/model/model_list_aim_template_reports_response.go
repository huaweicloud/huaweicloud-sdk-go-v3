package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAimTemplateReportsResponse Response Object
type ListAimTemplateReportsResponse struct {

	// 查询模板报表结果集。
	TemplateReports *[]AimTemplateReport `json:"template_reports,omitempty"`

	PageInfo       *Page `json:"page_info,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ListAimTemplateReportsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAimTemplateReportsResponse struct{}"
	}

	return strings.Join([]string{"ListAimTemplateReportsResponse", string(data)}, " ")
}
