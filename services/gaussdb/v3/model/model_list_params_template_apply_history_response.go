package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListParamsTemplateApplyHistoryResponse Response Object
type ListParamsTemplateApplyHistoryResponse struct {

	// 应用记录数量。
	TotalCount *string `json:"total_count,omitempty"`

	// 应用记录信息
	Histories      *[]TemplateApplyHistory `json:"histories,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListParamsTemplateApplyHistoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListParamsTemplateApplyHistoryResponse struct{}"
	}

	return strings.Join([]string{"ListParamsTemplateApplyHistoryResponse", string(data)}, " ")
}
