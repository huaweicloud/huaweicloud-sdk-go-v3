package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlertFileInfo struct {

	// 文件路径/名称
	FilePath *string `json:"file_path,omitempty"`

	// 文件内容
	FileContent *string `json:"file_content,omitempty"`

	// 文件新路径/名称
	FileNewPath *string `json:"file_new_path,omitempty"`

	// 文件hash
	FileHash *string `json:"file_hash,omitempty"`

	// 文件md5
	FileMd5 *string `json:"file_md5,omitempty"`

	// 文件sha256
	FileSha256 *string `json:"file_sha256,omitempty"`

	// 文件属性
	FileAttr *string `json:"file_attr,omitempty"`
}

func (o AlertFileInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlertFileInfo struct{}"
	}

	return strings.Join([]string{"AlertFileInfo", string(data)}, " ")
}
