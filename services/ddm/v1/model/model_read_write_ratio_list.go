package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReadWriteRatioList struct {

	// 逻辑库名称。
	Schema *string `json:"schema,omitempty" xml:"schema"`

	// 逻辑表名称。
	Table *string `json:"table,omitempty" xml:"table"`

	// 读次数。
	ReadCount *string `json:"readCount,omitempty" xml:"readCount"`

	// 写次数。
	WriteCount *string `json:"writeCount,omitempty" xml:"writeCount"`

	// 关联表。
	RelationTables *string `json:"relationTables,omitempty" xml:"relationTables"`

	// 最后执行时间。
	LastUpdated *string `json:"lastUpdated,omitempty" xml:"lastUpdated"`
}

func (o ReadWriteRatioList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReadWriteRatioList struct{}"
	}

	return strings.Join([]string{"ReadWriteRatioList", string(data)}, " ")
}
