package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KernelModuleInfo 内核模块信息
type KernelModuleInfo struct {

	// **参数解释**: 内核模块名称 **取值范围**: 字符长度0-256
	Name *string `json:"name,omitempty"`

	// **参数解释**: 文件名称 **取值范围**: 字符长度0-256
	FileName *string `json:"file_name,omitempty"`

	// **参数解释**: 内核模块版本号 **取值范围**: 字符长度0-64
	Version *string `json:"version,omitempty"`

	// **参数解释**: 源码版本号 **取值范围**: 字符长度0-64
	Srcversion *string `json:"srcversion,omitempty"`

	// **参数解释**: 文件路径 **取值范围**: 字符长度0-1024
	Path *string `json:"path,omitempty"`

	// **参数解释**: 文件大小 **取值范围**: 最小值0，最大值2147483647
	Size *int32 `json:"size,omitempty"`

	// **参数解释**: 文件权限 **取值范围**: 字符长度0-32
	Mode *string `json:"mode,omitempty"`

	// **参数解释**: 文件用户ID **取值范围**: 最小值0，最大值2147483647
	Uid *int32 `json:"uid,omitempty"`

	// **参数解释**: 文件创建时间 **取值范围**: 最小值0，最大值2^63-1
	Ctime *int64 `json:"ctime,omitempty"`

	// **参数解释**: 最后修改时间 **取值范围**: 最小值0，最大值2^63-1
	Mtime *int64 `json:"mtime,omitempty"`

	// **参数解释**: 文件哈希 **取值范围**: 字符长度0-64
	Hash *string `json:"hash,omitempty"`

	// **参数解释**: 内核模块描述信息 **取值范围**: 字符长度0-256
	Desc *string `json:"desc,omitempty"`

	// **参数解释**: 扫描时间 **取值范围**: 最小值0，最大值2^63-1
	RecordTime *int64 `json:"record_time,omitempty"`
}

func (o KernelModuleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KernelModuleInfo struct{}"
	}

	return strings.Join([]string{"KernelModuleInfo", string(data)}, " ")
}
