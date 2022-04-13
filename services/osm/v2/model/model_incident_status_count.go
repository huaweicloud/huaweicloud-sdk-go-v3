package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IncidentStatusCount struct {
	// 状态 0：待受理 1：处理中 2：待确认结果 3：已完成 4：已撤销 12：无效 17： 待反馈

	Status *int32 `json:"status,omitempty"`
	// 数量

	Count *int32 `json:"count,omitempty"`
}

func (o IncidentStatusCount) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IncidentStatusCount struct{}"
	}

	return strings.Join([]string{"IncidentStatusCount", string(data)}, " ")
}
