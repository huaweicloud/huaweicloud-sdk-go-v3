package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WebTamperProtectionDirectoryResponseInfo **参数解释**:  防护配置关联的容器的防护目录 **取值范围**: 不涉及
type WebTamperProtectionDirectoryResponseInfo struct {

	// **参数解释**： 集群ID **取值范围**： 字符长度1-256位
	ProtectDirectory *string `json:"protect_directory,omitempty"`

	// **参数解释**: 防护状态 **取值范围**: - protected：防护中 - protect_failed：防护失败
	Status *string `json:"status,omitempty"`

	// **参数解释**： 防护失败原因 **取值范围**： 字符长度1-256位
	FailedReason *string `json:"failed_reason,omitempty"`

	// **参数解释**： 本地备份路径 **取值范围**： 字符长度1-256位
	BackupDirectory *string `json:"backup_directory,omitempty"`

	// **参数解释**： 需要排除的子目录列表 **取值范围**： 最少0条，最多10条
	ExcludeChildDirectoryList *[]string `json:"exclude_child_directory_list,omitempty"`

	// **参数解释**： 需要排除的子文件列表 **取值范围**： 最少0条，最多10条
	ExcludeFileList *[]string `json:"exclude_file_list,omitempty"`
}

func (o WebTamperProtectionDirectoryResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebTamperProtectionDirectoryResponseInfo struct{}"
	}

	return strings.Join([]string{"WebTamperProtectionDirectoryResponseInfo", string(data)}, " ")
}
