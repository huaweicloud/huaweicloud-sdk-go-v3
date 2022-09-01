package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 数据行
type FrontRow struct {

	// 数据单元集合
	CellList *[]FrontCell `json:"cell_list,omitempty" xml:"cell_list"`

	// 将group by的字段拼接成过滤字符串，用于后续点网格点击使用
	Filter *string `json:"filter,omitempty" xml:"filter"`

	// 是否是header信息
	Header *bool `json:"header,omitempty" xml:"header"`

	// 是否是事务
	TxId *int64 `json:"tx_id,omitempty" xml:"tx_id"`
}

func (o FrontRow) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FrontRow struct{}"
	}

	return strings.Join([]string{"FrontRow", string(data)}, " ")
}
