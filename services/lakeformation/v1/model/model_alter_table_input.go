package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlterTableInput struct {

	// 修改表参数映射信息，支持的参数如下： CASADE: 级联删除列，如果为true则会把partition中的列也删除；如果为false则不会 DO_NOT_UPDATE_STATS: 不更新文件级别统计信息。true则不更新；false则更新。 STATS_GENERATED：记录本次更新的发起者。可填：TASK/USET。具体作用未明确。
	AlterParams map[string]string `json:"alter_params,omitempty"`

	Table *TableInput `json:"table"`
}

func (o AlterTableInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlterTableInput struct{}"
	}

	return strings.Join([]string{"AlterTableInput", string(data)}, " ")
}
