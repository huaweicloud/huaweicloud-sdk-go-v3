package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Statistic struct {

	// 用户名称
	UserName *string `json:"user_name,omitempty"`

	// 数据库名称
	Database *string `json:"database,omitempty"`

	// 由SQL的语法解析树计算出的内部哈希码。
	QueryId *int64 `json:"query_id,omitempty"`

	// 调用次数
	Calls *int64 `json:"calls,omitempty"`

	// SQL语句的文本形式。
	Query *string `json:"query,omitempty"`

	// 扫描行数
	Rows *int64 `json:"rows,omitempty"`

	// 是否可以执行sql限流
	CanUse *float64 `json:"can_use,omitempty"`
}

func (o Statistic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Statistic struct{}"
	}

	return strings.Join([]string{"Statistic", string(data)}, " ")
}
