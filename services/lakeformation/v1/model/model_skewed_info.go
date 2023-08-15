package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SkewedInfo struct {

	// 数据偏移列的列表
	SkewedColumnNames []string `json:"skewed_column_names"`

	// 偏斜值和地址的映射关系.
	SkewedColumnValueLocationMaps map[string]string `json:"skewed_column_value_location_maps"`

	// 偏斜值的列表.
	SkewedColumnValues [][]string `json:"skewed_column_values"`
}

func (o SkewedInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SkewedInfo struct{}"
	}

	return strings.Join([]string{"SkewedInfo", string(data)}, " ")
}
