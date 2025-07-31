package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KernelModuleInfo 内核模块信息
type KernelModuleInfo struct {

	// 内核模块名称
	Name *string `json:"name,omitempty"`

	// 文件名称
	FileName *string `json:"file_name,omitempty"`

	// 内核模块版本号
	Version *string `json:"version,omitempty"`

	// 源码版本号
	Srcversion *string `json:"srcversion,omitempty"`

	// 文件路径
	Path *string `json:"path,omitempty"`

	// 文件大小
	Size *int64 `json:"size,omitempty"`

	// 文件权限
	Mode *string `json:"mode,omitempty"`

	// 文件用户ID
	Uid *int32 `json:"uid,omitempty"`

	// 文件创建时间
	Ctime *int64 `json:"ctime,omitempty"`

	// 最后修改时间
	Mtime *int64 `json:"mtime,omitempty"`

	// 文件哈希
	Hash *string `json:"hash,omitempty"`

	// 内核模块描述信息
	Desc *string `json:"desc,omitempty"`

	// 扫描时间
	RecordTime *int64 `json:"record_time,omitempty"`

	// 首次扫描时间
	FirstScanTime *int64 `json:"first_scan_time,omitempty"`
}

func (o KernelModuleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KernelModuleInfo struct{}"
	}

	return strings.Join([]string{"KernelModuleInfo", string(data)}, " ")
}
