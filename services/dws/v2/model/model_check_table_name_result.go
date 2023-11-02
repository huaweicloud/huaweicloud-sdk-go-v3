package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckTableNameResult 表名检查结果
type CheckTableNameResult struct {

	// 数据库名称
	Database *string `json:"database,omitempty"`

	// 恢复源表信息
	RestoreTableList *[]string `json:"restore_table_list,omitempty"`

	// 恢复目的表信息
	TargetTableList *[]string `json:"target_table_list,omitempty"`
}

func (o CheckTableNameResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckTableNameResult struct{}"
	}

	return strings.Join([]string{"CheckTableNameResult", string(data)}, " ")
}
