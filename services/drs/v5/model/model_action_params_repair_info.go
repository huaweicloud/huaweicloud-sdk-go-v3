package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ActionParamsRepairInfo 数据修复信息。
type ActionParamsRepairInfo struct {

	// 对比任务ID。
	QueryId *string `json:"query_id,omitempty"`

	// 数据修复对象信息。
	Objects *[]ActionParamsRepairInfoObjects `json:"objects,omitempty"`
}

func (o ActionParamsRepairInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionParamsRepairInfo struct{}"
	}

	return strings.Join([]string{"ActionParamsRepairInfo", string(data)}, " ")
}
