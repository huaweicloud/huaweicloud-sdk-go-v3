package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BackupInfoResp 备份文件信息
type BackupInfoResp struct {

	// 备份文件来源，包括OBS和RDS两种。
	FileSource *string `json:"file_source,omitempty"`

	// OBS桶名称。
	BucketName *string `json:"bucket_name,omitempty"`

	// 备份文件列表。
	FileInfo *[]BackupFileResp `json:"file_info,omitempty"`
}

func (o BackupInfoResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupInfoResp struct{}"
	}

	return strings.Join([]string{"BackupInfoResp", string(data)}, " ")
}
