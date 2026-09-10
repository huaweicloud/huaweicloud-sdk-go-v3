package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShareBackupDatastore struct {

	// 数据库引擎，支持的引擎如下，不区分大小写：PostgreSQL。
	Type *string `json:"type,omitempty"`

	// 数据库版本。
	Version *string `json:"version,omitempty"`
}

func (o ShareBackupDatastore) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShareBackupDatastore struct{}"
	}

	return strings.Join([]string{"ShareBackupDatastore", string(data)}, " ")
}
