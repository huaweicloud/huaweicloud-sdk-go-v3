package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 备份实例请求体
type BackupInstanceBody struct {
	// 备份缓存实例的备注信息。

	Remark *string `json:"remark,omitempty"`
}

func (o BackupInstanceBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupInstanceBody struct{}"
	}

	return strings.Join([]string{"BackupInstanceBody", string(data)}, " ")
}
