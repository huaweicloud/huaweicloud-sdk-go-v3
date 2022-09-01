package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 文件信息
type EventFileResponseInfo struct {

	// 文件路径
	FilePath *string `json:"file_path,omitempty" xml:"file_path"`

	// 文件别名
	FileAlias *string `json:"file_alias,omitempty" xml:"file_alias"`

	// 文件大小
	FileSize *int32 `json:"file_size,omitempty" xml:"file_size"`

	// 文件最后一次修改时间
	FileMtime *int64 `json:"file_mtime,omitempty" xml:"file_mtime"`

	// 文件最后一次访问时间
	FileAtime *int64 `json:"file_atime,omitempty" xml:"file_atime"`

	// 文件最后一次状态改变时间
	FileCtime *int64 `json:"file_ctime,omitempty" xml:"file_ctime"`

	// 文件hash
	FileHash *string `json:"file_hash,omitempty" xml:"file_hash"`

	// 文件md5
	FileMd5 *string `json:"file_md5,omitempty" xml:"file_md5"`

	// 文件sha256
	FileSha256 *string `json:"file_sha256,omitempty" xml:"file_sha256"`

	// 文件类型
	FileType *string `json:"file_type,omitempty" xml:"file_type"`

	// 文件内容
	FileContent *string `json:"file_content,omitempty" xml:"file_content"`

	// 文件属性
	FileAttr *string `json:"file_attr,omitempty" xml:"file_attr"`

	// 文件操作类型
	FileOperation *int32 `json:"file_operation,omitempty" xml:"file_operation"`

	// 文件动作
	FileAction *string `json:"file_action,omitempty" xml:"file_action"`

	// 变更前后的属性
	FileChangeAttr *string `json:"file_change_attr,omitempty" xml:"file_change_attr"`

	// 新文件路径
	FileNewPath *string `json:"file_new_path,omitempty" xml:"file_new_path"`

	// 文件描述
	FileDesc *string `json:"file_desc,omitempty" xml:"file_desc"`

	// 文件关键字
	FileKeyWord *string `json:"file_key_word,omitempty" xml:"file_key_word"`

	// 是否目录
	IsDir *bool `json:"is_dir,omitempty" xml:"is_dir"`

	// 文件句柄信息
	FdInfo *string `json:"fd_info,omitempty" xml:"fd_info"`

	// 文件句柄数量
	FdCount *int32 `json:"fd_count,omitempty" xml:"fd_count"`
}

func (o EventFileResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventFileResponseInfo struct{}"
	}

	return strings.Join([]string{"EventFileResponseInfo", string(data)}, " ")
}
