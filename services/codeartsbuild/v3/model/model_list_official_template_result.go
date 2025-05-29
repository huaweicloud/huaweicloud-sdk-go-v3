package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOfficialTemplateResult 返回值
type ListOfficialTemplateResult struct {

	// 总数
	TotalSize float32 `json:"total_size,omitempty"`

	// 模版列表
	Items *[]TemplateList `json:"items,omitempty"`
}

func (o ListOfficialTemplateResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOfficialTemplateResult struct{}"
	}

	return strings.Join([]string{"ListOfficialTemplateResult", string(data)}, " ")
}
