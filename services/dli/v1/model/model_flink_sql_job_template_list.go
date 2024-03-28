package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FlinkSqlJobTemplateList struct {

	// 模板总数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 模板详细信息
	Templates *[]FlinkSqlJobTemplate `json:"templates,omitempty"`
}

func (o FlinkSqlJobTemplateList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkSqlJobTemplateList struct{}"
	}

	return strings.Join([]string{"FlinkSqlJobTemplateList", string(data)}, " ")
}
