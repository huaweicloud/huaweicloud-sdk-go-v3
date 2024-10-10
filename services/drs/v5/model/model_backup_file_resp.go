package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BackupFileResp 备份文件信息体。
type BackupFileResp struct {

	// 文件名称。
	FileName *string `json:"file_name,omitempty"`

	// 备份文件大小。
	FileSize *string `json:"file_size,omitempty"`

	// 备份文件最近修改时间。
	FileLastModify *string `json:"file_last_modify,omitempty"`
}

func (o BackupFileResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupFileResp struct{}"
	}

	return strings.Join([]string{"BackupFileResp", string(data)}, " ")
}
