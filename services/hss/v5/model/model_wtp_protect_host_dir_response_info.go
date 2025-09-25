package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WtpProtectHostDirResponseInfo struct {

	// **参数解释**: 防护目录 **取值范围**: 字符长度0-512位
	ProtectDir *string `json:"protect_dir,omitempty"`

	// **参数解释**: 排除子目录 **取值范围**: 字符长度0-512位
	ExcludeChildDir *string `json:"exclude_child_dir,omitempty"`

	// 排除文件路径
	ExclueFilePath *string `json:"exclue_file_path,omitempty"`

	// **参数解释**: 本地备份路径，仅Linux服务器支持设置本地备份路径。 **取值范围**: 字符长度0-512位
	LocalBackupDir *string `json:"local_backup_dir,omitempty"`

	// **参数解释**: 防护状态 **取值范围**: - closed ：未开启。 - opened ：防护中。 - opening ：开启中。 - closing ：关闭中。 - open_failed ：防护失败。
	ProtectStatus *string `json:"protect_status,omitempty"`

	// **参数解释**: 失败原因，当防护状态为open_failed时存在失败原因 **取值范围**: 字符长度0-512位
	Error *string `json:"error,omitempty"`
}

func (o WtpProtectHostDirResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WtpProtectHostDirResponseInfo struct{}"
	}

	return strings.Join([]string{"WtpProtectHostDirResponseInfo", string(data)}, " ")
}
