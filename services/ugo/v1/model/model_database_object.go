package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 语法转换或者迁移的数据库对象。
type DatabaseObject struct {

	// 对象类型。
	ObjectType string `json:"object_type"`

	// 该类型对象的总数。
	TotalCount int64 `json:"total_count"`

	// 成功的对象数量。
	SucceedCount int64 `json:"succeed_count"`

	// 失败的对象数量。
	FailedCount int64 `json:"failed_count"`

	// 忽略的对象数量。
	IgnoredCount int64 `json:"ignored_count"`

	// 手动操作的对象数量。
	ManualCount int64 `json:"manual_count"`

	// 成功率。
	SuccessRate string `json:"success_rate"`
}

func (o DatabaseObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabaseObject struct{}"
	}

	return strings.Join([]string{"DatabaseObject", string(data)}, " ")
}
