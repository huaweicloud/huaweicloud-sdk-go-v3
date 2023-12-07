package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CoverRateVo 需求覆盖率
type CoverRateVo struct {

	// 统计测试中的需求
	Testing *int32 `json:"testing,omitempty"`

	// 统计已完成的需求
	Finished *int32 `json:"finished,omitempty"`

	// 统计未测试的需求
	NotTested *int32 `json:"not_tested,omitempty"`

	// 计算需求总数
	TotalNumber *int32 `json:"total_number,omitempty"`

	// 需求覆盖率
	CoverRate *string `json:"cover_rate,omitempty"`
}

func (o CoverRateVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CoverRateVo struct{}"
	}

	return strings.Join([]string{"CoverRateVo", string(data)}, " ")
}
