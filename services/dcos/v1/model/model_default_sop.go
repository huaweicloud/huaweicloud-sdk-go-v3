package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DefaultSop 默认的附件sop
type DefaultSop struct {

	// 中文版本文件名称
	FileCnName *string `json:"file_cn_name,omitempty"`

	// 中文版本文件id
	FileCnId *string `json:"file_cn_id,omitempty"`

	// 英文版本文件名称
	FileEnName *string `json:"file_en_name,omitempty"`

	// 英文版本文件id
	FileEnId *string `json:"file_en_id,omitempty"`

	// 归属组织
	Ownership *string `json:"ownership,omitempty"`

	// 是否默认
	DefaultOptions bool `json:"default_options"`
}

func (o DefaultSop) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DefaultSop struct{}"
	}

	return strings.Join([]string{"DefaultSop", string(data)}, " ")
}
