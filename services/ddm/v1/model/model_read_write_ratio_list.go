package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReadWriteRatioList struct {

	// 逻辑库名称。
	Schema *string `json:"schema,omitempty"`

	// 逻辑表名称。
	Table *string `json:"table,omitempty"`

	// 读次数。
	ReadCount *string `json:"readCount,omitempty"`

	// 写次数。
	WriteCount *string `json:"writeCount,omitempty"`

	// 关联表。
	RelationTables *string `json:"relationTables,omitempty"`

	// 最后执行时间。
	LastUpdated *string `json:"lastUpdated,omitempty"`
}

func (o ReadWriteRatioList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReadWriteRatioList struct{}"
	}

	return strings.Join([]string{"ReadWriteRatioList", string(data)}, " ")
}
