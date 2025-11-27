package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProtectDirectoryInfo **参数解释**: 防护目录信息 **约束限制**: 不涉及 **取值范围**: 不涉及 **默认取值**: 不涉及
type ProtectDirectoryInfo struct {

	// **参数解释**： 防护目录 **约束限制**: 不涉及 **取值范围**： 字符长度1-256位 **默认取值**: 不涉及
	ProtectDirectory string `json:"protect_directory"`

	// **参数解释**： 需要排除的子目录列表 **约束限制**: 不涉及 **取值范围**： 最少0条，最多10条 **默认取值**: 不涉及
	ExcludeChildDirectoryList *[]string `json:"exclude_child_directory_list,omitempty"`

	// **参数解释**： 需要排除的子文件列表 **约束限制**: 不涉及 **取值范围**： 最少0条，最多10条 **默认取值**: 不涉及
	ExcludeFileList *[]string `json:"exclude_file_list,omitempty"`

	// **参数解释**： 备份目录 **约束限制**: 不涉及 **取值范围**： 字符长度1-256位 **默认取值**: 不涉及
	BackupDirectory *string `json:"backup_directory,omitempty"`
}

func (o ProtectDirectoryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectDirectoryInfo struct{}"
	}

	return strings.Join([]string{"ProtectDirectoryInfo", string(data)}, " ")
}
