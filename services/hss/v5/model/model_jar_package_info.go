package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JarPackageInfo **参数解释**: 中间件信息
type JarPackageInfo struct {

	// **参数解释**: 中间件名称 **取值范围**: 字符长度1-256
	FileName *string `json:"file_name,omitempty"`

	// **参数解释**: 中间件类型 **取值范围**: -jar：jar包 -python：python软件包 -npm：npm软件包
	FileType *string `json:"file_type,omitempty"`

	// **参数解释**: 分类 **取值范围**: -util：默认分类
	Catalogue *string `json:"catalogue,omitempty"`

	// **参数解释**: 中间件版本 **取值范围**: 字符长度1-512
	Version *string `json:"version,omitempty"`

	// **参数解释**: 中间件路径 **取值范围**: 字符长度1-512
	Path *string `json:"path,omitempty"`

	// **参数解释**: 中间件文件权限 **取值范围**: 字符长度1-32
	Mode *string `json:"mode,omitempty"`

	// **参数解释**: uid **取值范围**: 最小值0，最大值2147483647
	Uid *int32 `json:"uid,omitempty"`

	// **参数解释**: 中间件文件hash **取值范围**: 字符长度1-64
	Hash *string `json:"hash,omitempty"`

	// **参数解释**: 中间件进程ID **取值范围**: 最小值0，最大值2147483647
	Pid *int32 `json:"pid,omitempty"`

	// **参数解释**: 中间件进程路径 **取值范围**: 字符长度1-1024
	ProcPath *string `json:"proc_path,omitempty"`

	// **参数解释**: 中间件扫描时间 **取值范围**: 最小值0，最大值2^63-1
	RecordTime *int64 `json:"record_time,omitempty"`

	// **参数解释**: 容器id **取值范围**: 字符长度1-128
	ContainerId *string `json:"container_id,omitempty"`

	// **参数解释**: 容器名称 **取值范围**: 字符长度1-256
	ContainerName *string `json:"container_name,omitempty"`
}

func (o JarPackageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JarPackageInfo struct{}"
	}

	return strings.Join([]string{"JarPackageInfo", string(data)}, " ")
}
