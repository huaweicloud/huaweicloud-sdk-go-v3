package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResultFileResponseInfo 文件信息
type ResultFileResponseInfo struct {

	// 文件路径
	FilePath *string `json:"file_path,omitempty"`

	// 文件哈希
	FileHash *string `json:"file_hash,omitempty"`

	// 文件大小
	FileSize *int64 `json:"file_size,omitempty"`

	// 文件属主
	FileOwner *string `json:"file_owner,omitempty"`

	// 文件属性
	FileAttr *string `json:"file_attr,omitempty"`

	// 文件创建时间
	FileCtime *int64 `json:"file_ctime,omitempty"`

	// 文件更新时间
	FileMtime *int64 `json:"file_mtime,omitempty"`
}

func (o ResultFileResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResultFileResponseInfo struct{}"
	}

	return strings.Join([]string{"ResultFileResponseInfo", string(data)}, " ")
}
