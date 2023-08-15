package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlinkTemplateList
type FlinkTemplateList struct {

	// 模板总数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 模板详细信息
	Templates *[]FlinkTemplateDetail `json:"templates,omitempty"`
}

func (o FlinkTemplateList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkTemplateList struct{}"
	}

	return strings.Join([]string{"FlinkTemplateList", string(data)}, " ")
}
