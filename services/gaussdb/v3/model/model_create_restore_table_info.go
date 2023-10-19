package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateRestoreTableInfo struct {

	// 恢复前表名。
	OldName string `json:"old_name"`

	// 恢复后表名。
	NewName string `json:"new_name"`
}

func (o CreateRestoreTableInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRestoreTableInfo struct{}"
	}

	return strings.Join([]string{"CreateRestoreTableInfo", string(data)}, " ")
}
