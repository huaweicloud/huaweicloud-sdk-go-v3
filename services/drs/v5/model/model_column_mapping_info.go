package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ColumnMappingInfo 列映射信息
type ColumnMappingInfo struct {

	// 对象id
	ObjectId *string `json:"object_id,omitempty"`

	// 列信息
	ColumnInfoLists *[]ColumnInfo `json:"column_info_lists,omitempty"`
}

func (o ColumnMappingInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ColumnMappingInfo struct{}"
	}

	return strings.Join([]string{"ColumnMappingInfo", string(data)}, " ")
}
