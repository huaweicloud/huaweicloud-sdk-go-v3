package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TopSqlsRowResult struct {

	// 对查询计算的二进制哈希值，用于标识具有类似逻辑的查询。
	Id *string `json:"id,omitempty"`

	// 数据类型。 取值范围： AvgReturnRows 平均返回行 TotalReturnRows 总返回行
	DataType *string `json:"data_type,omitempty"`

	// 行数。
	Value *string `json:"value,omitempty"`
}

func (o TopSqlsRowResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopSqlsRowResult struct{}"
	}

	return strings.Join([]string{"TopSqlsRowResult", string(data)}, " ")
}
