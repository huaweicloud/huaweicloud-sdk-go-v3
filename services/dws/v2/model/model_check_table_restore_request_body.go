package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckTableRestoreRequestBody 恢复表前检查表名请求体
type CheckTableRestoreRequestBody struct {

	// 名称是否区分大小写
	CaseSensitive bool `json:"case_sensitive"`

	// 数据库名称
	Database string `json:"database"`

	// 源表信息
	RestoreTableList []TableDetail `json:"restore_table_list"`

	// 目的表信息
	TargetTableList []TableDetail `json:"target_table_list"`
}

func (o CheckTableRestoreRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckTableRestoreRequestBody struct{}"
	}

	return strings.Join([]string{"CheckTableRestoreRequestBody", string(data)}, " ")
}
