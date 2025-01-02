package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DelayRestoreSchema struct {

	// schema名称
	Name string `json:"name"`

	// 返回该schema下的总表数量
	TotalTables *int32 `json:"total_tables,omitempty"`

	// 该schema下的表列表
	Tables *[]DelayRestoreTable `json:"tables,omitempty"`
}

func (o DelayRestoreSchema) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DelayRestoreSchema struct{}"
	}

	return strings.Join([]string{"DelayRestoreSchema", string(data)}, " ")
}
