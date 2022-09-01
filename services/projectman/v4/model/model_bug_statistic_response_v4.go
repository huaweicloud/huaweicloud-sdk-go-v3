package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Bug信息
type BugStatisticResponseV4 struct {

	// 重要程度为关键的缺陷数
	CriticalNum *int32 `json:"critical_num,omitempty" xml:"critical_num"`

	// DI
	DefectIndex *float64 `json:"defect_index,omitempty" xml:"defect_index"`

	// 模块
	Module *string `json:"module,omitempty" xml:"module"`

	// 重要程度为一般的缺陷数
	NormalNum *int32 `json:"normal_num,omitempty" xml:"normal_num"`

	// 重要程度为严重的缺陷数
	SeriousNum *int32 `json:"serious_num,omitempty" xml:"serious_num"`

	// 重要程度为提示的缺陷数
	TipNum *int32 `json:"tip_num,omitempty" xml:"tip_num"`

	// 总数
	Total *int32 `json:"total,omitempty" xml:"total"`
}

func (o BugStatisticResponseV4) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BugStatisticResponseV4 struct{}"
	}

	return strings.Join([]string{"BugStatisticResponseV4", string(data)}, " ")
}
