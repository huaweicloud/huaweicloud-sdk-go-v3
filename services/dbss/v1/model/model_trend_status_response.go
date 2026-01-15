package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TrendStatusResponse 数据库趋势统计
type TrendStatusResponse struct {

	// 数据库ID
	DbId *string `json:"db_id,omitempty"`

	// 数据库名称
	DbName *string `json:"db_name,omitempty"`

	// 生成时间
	GenerateTime *string `json:"generate_time,omitempty"`

	// 统计类型  - RISK：风险  - SESSION：SESSION  - SQL：SQL  - COUNT：数量  - INJECTION：注入 - OPERATION：操作
	StatisticsType *string `json:"statistics_type,omitempty"`
}

func (o TrendStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrendStatusResponse struct{}"
	}

	return strings.Join([]string{"TrendStatusResponse", string(data)}, " ")
}
