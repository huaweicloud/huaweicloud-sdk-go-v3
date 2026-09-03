package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RowPairDto 变更数据详情
type RowPairDto struct {

	// 变更前的行数据
	BeforeRow *interface{} `json:"before_row,omitempty"`

	// 变更后的行数据
	AfterRow *interface{} `json:"after_row,omitempty"`
}

func (o RowPairDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RowPairDto struct{}"
	}

	return strings.Join([]string{"RowPairDto", string(data)}, " ")
}
