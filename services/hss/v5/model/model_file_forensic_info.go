package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FileForensicInfo 文件取证信息
type FileForensicInfo struct {

	// **参数解释**： 文件路径 **取值范围**： 不涉及
	FilePath *string `json:"file_path,omitempty"`

	// **参数解释**： 文件新路径 **取值范围**： 不涉及
	FileNewPath *string `json:"file_new_path,omitempty"`

	// **参数解释**： 文件名称 **取值范围**： 不涉及
	FileName *string `json:"file_name,omitempty"`

	// **参数解释**： 文件sha256 **取值范围**： 不涉及
	FileSha256 *string `json:"file_sha256,omitempty"`

	// **参数解释**： 文件动作 **取值范围**： 不涉及
	FileAction *string `json:"file_action,omitempty"`

	// **参数解释**： 文件操作类型 **取值范围**： 不涉及
	FileOperation *int32 `json:"file_operation,omitempty"`

	// **参数解释**： 文件大小 **取值范围**： 不涉及
	FileSize *int32 `json:"file_size,omitempty"`

	// **参数解释**： 文件hash,当前为sha256 **取值范围**： 不涉及
	FileHash *string `json:"file_hash,omitempty"`

	// **参数解释**： 文件描述 **取值范围**： 不涉及
	FileDesc *string `json:"file_desc,omitempty"`

	// **参数解释**： 是否目录 **取值范围**： 不涉及
	IsDir *bool `json:"is_dir,omitempty"`

	// **参数解释**： 文件最后一次修改时间(毫秒) **取值范围**： 不涉及
	FileMtime *int64 `json:"file_mtime,omitempty"`

	// **参数解释**： 文件最后一次访问时间(毫秒) **取值范围**： 不涉及
	FileAtime *int64 `json:"file_atime,omitempty"`

	// **参数解释**： 文件最后一次状态改变时间(毫秒) **取值范围**： 不涉及
	FileCtime *int64 `json:"file_ctime,omitempty"`

	// **参数解释**： 文件别名 **取值范围**： 不涉及
	FileAlias *string `json:"file_alias,omitempty"`

	// **参数解释**： 文件md5 **取值范围**： 不涉及
	FileMd5 *string `json:"file_md5,omitempty"`

	// **参数解释**： 文件类型 **取值范围**： 不涉及
	FileType *string `json:"file_type,omitempty"`

	// **参数解释**： 文件关键字 **取值范围**： 不涉及
	FileKeyWord *string `json:"file_key_word,omitempty"`
}

func (o FileForensicInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FileForensicInfo struct{}"
	}

	return strings.Join([]string{"FileForensicInfo", string(data)}, " ")
}
