package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResultFileResponseInfo 文件信息
type ResultFileResponseInfo struct {

	// **参数解释**： 文件路径 **取值范围**： 字符长度1-256位
	FilePath *string `json:"file_path,omitempty"`

	// **参数解释**： 文件哈希 **取值范围**： 字符长度1-256位
	FileHash *string `json:"file_hash,omitempty"`

	// **参数解释**: 文件大小 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值9223372036854775807 **默认取值**: 不涉及
	FileSize *int64 `json:"file_size,omitempty"`

	// **参数解释**： 文件属性 **取值范围**： 字符长度0-64位
	FileOwner *string `json:"file_owner,omitempty"`

	// **参数解释**： 文件的系统属性（如读写权限、隐藏属性、执行权限等） **取值范围**： 字符长度1-256位
	FileAttr *string `json:"file_attr,omitempty"`

	// **参数解释**： 文件创建时间 **取值范围**： 非负长整数，时间格式：毫秒级时间戳（UTC时区，从1970-01-01 00:00:00开始计算），单位：ms
	FileCtime *int64 `json:"file_ctime,omitempty"`

	// **参数解释**： 文件更新时间 **取值范围**： 非负长整数，时间格式：毫秒级时间戳（UTC时区，从1970-01-01 00:00:00开始计算），单位：ms
	FileMtime *int64 `json:"file_mtime,omitempty"`
}

func (o ResultFileResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResultFileResponseInfo struct{}"
	}

	return strings.Join([]string{"ResultFileResponseInfo", string(data)}, " ")
}
