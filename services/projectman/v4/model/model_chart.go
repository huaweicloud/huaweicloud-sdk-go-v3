package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Chart struct {

	// 统计时间
	Date *string `json:"date,omitempty" xml:"date"`

	// 完成story工单
	FinishedNum *int32 `json:"finished_num,omitempty" xml:"finished_num"`

	// 迭代id
	IterationId *int32 `json:"iteration_id,omitempty" xml:"iteration_id"`

	// 项目id
	ProjectNumId *int32 `json:"project_num_id,omitempty" xml:"project_num_id"`

	// 未完成story数
	RemainingNum *int32 `json:"remaining_num,omitempty" xml:"remaining_num"`

	// 总story数
	Total *int32 `json:"total,omitempty" xml:"total"`
}

func (o Chart) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Chart struct{}"
	}

	return strings.Join([]string{"Chart", string(data)}, " ")
}
