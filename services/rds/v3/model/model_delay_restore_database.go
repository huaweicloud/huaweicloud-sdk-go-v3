package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DelayRestoreDatabase struct {

	// 数据库名称
	Name string `json:"name"`

	// 返回该库下的总表数量
	TotalTables *int32 `json:"total_tables,omitempty"`

	// 该库下的schema列表
	Schemas *[]DelayRestoreSchema `json:"schemas,omitempty"`
}

func (o DelayRestoreDatabase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DelayRestoreDatabase struct{}"
	}

	return strings.Join([]string{"DelayRestoreDatabase", string(data)}, " ")
}
