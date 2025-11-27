package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProtectDirectoryResponseInfo **参数解释**: 防护目录信息 **取值范围**: 不涉及
type ProtectDirectoryResponseInfo struct {

	// **参数解释**： 防护目录 **取值范围**： 字符长度1-256位
	ProtectDirectory *string `json:"protect_directory,omitempty"`

	// **参数解释**： 需要排除的子目录列表 **取值范围**： 最少0条，最多10条
	ExcludeChildDirectoryList *[]string `json:"exclude_child_directory_list,omitempty"`

	// **参数解释**： 需要排除的子文件列表 **取值范围**： 最少0条，最多10条
	ExcludeFileList *[]string `json:"exclude_file_list,omitempty"`

	// **参数解释**： 备份目录 **取值范围**： 字符长度1-256位
	BackupDirectory *string `json:"backup_directory,omitempty"`
}

func (o ProtectDirectoryResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectDirectoryResponseInfo struct{}"
	}

	return strings.Join([]string{"ProtectDirectoryResponseInfo", string(data)}, " ")
}
