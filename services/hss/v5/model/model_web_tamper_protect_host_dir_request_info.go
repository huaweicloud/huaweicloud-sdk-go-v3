package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WebTamperProtectHostDirRequestInfo 防护目录信息
type WebTamperProtectHostDirRequestInfo struct {

	// **参数解释**: 防护目录 **约束限制**: 不涉及 **取值范围**: 字符长度1-256位。Linux服务器，必须以/开头，不能以/结尾，只能包含英文大小写字母，数字，下划线，中划线和点。Windows服务器，目录名不能包含;/_*?\"<>|，不能以空格开头，不能以\\结尾。 **默认取值**: 不涉及
	ProtectDir string `json:"protect_dir"`

	// **参数解释**: 排除子目录 **约束限制**: 不涉及 **取值范围**: 子目录名必须是防护目录的有效相对路径，目录名最大长度不能超过256个字符，最多可添加10个子目录，多个子目录用;隔开。Linux服务器的子目录名不能以/开头或结尾，Windows服务器的子目录名不能以\\开头或结尾。 **默认取值**: 不涉及
	ExcludeChildDir *string `json:"exclude_child_dir,omitempty"`

	// **参数解释**: 排除文件路径 **约束限制**: 仅Linux服务器支持填写排除文件路径，Windows服务器不可填写该字段。 **取值范围**: 排除文件路径必须是防护目录的有效相对路径，不能以/开头或结尾，文件路径最大长度不能超过256个字符；最多可添加50个文件路径，多个文件路径用;隔开。 **默认取值**: 不涉及
	ExcludeFilePath *string `json:"exclude_file_path,omitempty"`

	// **参数解释**: 本地备份路径，Linux服务器必须填写该字段。 **约束限制**: 仅Linux服务器需要填写本地备份路径，Windows服务器不可填写该字段。 **取值范围**: 本地备份路径不能包含;字符，不能以空格开头，不能以/结尾，本地备份路径长度不得超过256个字符。 **默认取值**: 不涉及
	LocalBackupDir *string `json:"local_backup_dir,omitempty"`
}

func (o WebTamperProtectHostDirRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebTamperProtectHostDirRequestInfo struct{}"
	}

	return strings.Join([]string{"WebTamperProtectHostDirRequestInfo", string(data)}, " ")
}
