package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StatisticsRvo struct {

	// 新增统计记录数。
	CreateCount *int32 `json:"createCount,omitempty"`

	// 删除统计记录数。
	DeleteCount *int32 `json:"deleteCount,omitempty"`

	// 软删除统计记录数。
	LogicalDeleteCount *int32 `json:"logicalDeleteCount,omitempty"`

	// 更新统计记录数。
	UpdateCount *int32 `json:"updateCount,omitempty"`
}

func (o StatisticsRvo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatisticsRvo struct{}"
	}

	return strings.Join([]string{"StatisticsRvo", string(data)}, " ")
}
