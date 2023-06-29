package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryInstanceBackupResponseBodyDatastore 数据库信息。
type QueryInstanceBackupResponseBodyDatastore struct {

	// 数据库类型。
	Type string `json:"type"`

	// 数据库版本。
	Version string `json:"version"`
}

func (o QueryInstanceBackupResponseBodyDatastore) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryInstanceBackupResponseBodyDatastore struct{}"
	}

	return strings.Join([]string{"QueryInstanceBackupResponseBodyDatastore", string(data)}, " ")
}
