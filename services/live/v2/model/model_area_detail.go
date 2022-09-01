package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AreaDetail struct {

	// 各个计费大区名称，例如CN
	Area string `json:"area" xml:"area"`

	// 时间戳及相应时间的指标数值
	Summary []TimeValue `json:"summary" xml:"summary"`

	// 各个大区下的具体省份、区域、国家的时间戳及相应时间的指标数值
	Detail []AreaTimeValue `json:"detail" xml:"detail"`
}

func (o AreaDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AreaDetail struct{}"
	}

	return strings.Join([]string{"AreaDetail", string(data)}, " ")
}
