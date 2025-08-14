package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventFileResponseInfo 文件信息
type EventFileResponseInfo struct {

	// **参数解释**： 文件路径 **取值范围**： 字符长度1-256位
	FilePath *string `json:"file_path,omitempty"`

	// **参数解释**： 文件别名 **取值范围**： 字符长度1-256位
	FileAlias *string `json:"file_alias,omitempty"`

	// **参数解释**： 文件大小 **取值范围**： 最小值0，最大值2147483647
	FileSize *int32 `json:"file_size,omitempty"`

	// **参数解释**： 文件最后一次修改时间 **取值范围**： 最小值0，最大值9223372036854775807
	FileMtime *int64 `json:"file_mtime,omitempty"`

	// **参数解释**： 文件最后一次访问时间 **取值范围**： 最小值0，最大值9223372036854775807
	FileAtime *int64 `json:"file_atime,omitempty"`

	// **参数解释**： 文件最后一次状态改变时间 **取值范围**： 最小值0，最大值9223372036854775807
	FileCtime *int64 `json:"file_ctime,omitempty"`

	// **参数解释**： 文件hash,当前为sha256 **取值范围**： 字符长度1-256位
	FileHash *string `json:"file_hash,omitempty"`

	// **参数解释**： 文件md5 **取值范围**： 字符长度1-256位
	FileMd5 *string `json:"file_md5,omitempty"`

	// **参数解释**： 文件sha256 **取值范围**： 字符长度1-256位
	FileSha256 *string `json:"file_sha256,omitempty"`

	// **参数解释**： 文件类型 **取值范围**： 字符长度1-256位
	FileType *string `json:"file_type,omitempty"`

	// **参数解释**： 文件内容 **取值范围**： 字符长度1-256位
	FileContent *string `json:"file_content,omitempty"`

	// **参数解释**： 文件属性 **取值范围**： 字符长度1-256位
	FileAttr *string `json:"file_attr,omitempty"`

	// **参数解释**： 文件操作类型 **取值范围**： 最小值0，最大值2147483647
	FileOperation *int32 `json:"file_operation,omitempty"`

	// **参数解释**： 文件动作 **取值范围**： 字符长度1-256位
	FileAction *string `json:"file_action,omitempty"`

	// **参数解释**： 变更前后的属性 **取值范围**： 字符长度1-256位
	FileChangeAttr *string `json:"file_change_attr,omitempty"`

	// **参数解释**： 新文件路径 **取值范围**： 字符长度1-256位
	FileNewPath *string `json:"file_new_path,omitempty"`

	// **参数解释**： 文件描述 **取值范围**： 字符长度1-256位
	FileDesc *string `json:"file_desc,omitempty"`

	// **参数解释**： 文件关键字 **取值范围**： 字符长度1-256位
	FileKeyWord *string `json:"file_key_word,omitempty"`

	// **参数解释**： 是否目录 **取值范围**： - true：是目录 - false：不是目录
	IsDir *bool `json:"is_dir,omitempty"`

	// **参数解释**： 文件句柄信息 **取值范围**： 字符长度1-256位
	FdInfo *string `json:"fd_info,omitempty"`

	// **参数解释**： 文件句柄数量 **取值范围**： 最小值0，最大值2147483647
	FdCount *int32 `json:"fd_count,omitempty"`
}

func (o EventFileResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventFileResponseInfo struct{}"
	}

	return strings.Join([]string{"EventFileResponseInfo", string(data)}, " ")
}
