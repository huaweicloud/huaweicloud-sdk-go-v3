package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlinkJobList 作业列表详情。
type FlinkJobList struct {

	// 参数解释:  作业查询结果条数 示例: 1 约束限制:  无 取值范围: 大于等于0的整数 默认取值: 无
	TotalCount *int64 `json:"total_count,omitempty"`

	// 作业信息
	Jobs *[]FlinkJob `json:"jobs,omitempty"`
}

func (o FlinkJobList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkJobList struct{}"
	}

	return strings.Join([]string{"FlinkJobList", string(data)}, " ")
}
