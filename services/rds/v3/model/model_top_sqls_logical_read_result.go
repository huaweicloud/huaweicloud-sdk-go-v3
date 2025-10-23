package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TopSqlsLogicalReadResult struct {

	// 对查询计算的二进制哈希值，用于标识具有类似逻辑的查询。
	Id *string `json:"id,omitempty"`

	// 数据类型。取值范围： AvgLogicalReads 平均逻辑读 TotalLogicalReads 总逻辑读
	DataType *string `json:"data_type,omitempty"`

	// 逻辑读消耗。
	Value *string `json:"value,omitempty"`
}

func (o TopSqlsLogicalReadResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopSqlsLogicalReadResult struct{}"
	}

	return strings.Join([]string{"TopSqlsLogicalReadResult", string(data)}, " ")
}
