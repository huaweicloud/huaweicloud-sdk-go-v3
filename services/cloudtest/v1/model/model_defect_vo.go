package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DefectVo 整体缺陷信息
type DefectVo struct {

	// 缺陷数
	Total *int32 `json:"total,omitempty"`

	// 未关闭缺陷数
	NotSolved *int32 `json:"not_solved,omitempty"`

	// 组装缺陷每种重要程度的名称和对应的数目
	SeverityNumberList *[]NameAndValueVo `json:"severity_number_list,omitempty"`
}

func (o DefectVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DefectVo struct{}"
	}

	return strings.Join([]string{"DefectVo", string(data)}, " ")
}
