package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FileStatusV2 struct {

	// 文件在当前目录下的后缀，如获取“/tmp”目录，下面的“/tmp/test”文件，此处path_suffix内容为“test”。
	PathSuffix *string `json:"path_suffix,omitempty"`

	// 文件拥有者。
	Owner *string `json:"owner,omitempty"`

	// 文件属组。
	Group *string `json:"group,omitempty"`

	// 权限信息。
	Permission *string `json:"permission,omitempty"`

	// 副本数。
	Replication *int32 `json:"replication,omitempty"`

	// 块大小。
	BlockSize *int32 `json:"block_size,omitempty"`

	// 文件长度。
	Length *int32 `json:"length,omitempty"`

	// 文件类型：  - FILE：文件  - DIRECTORY：目录
	Type *string `json:"type,omitempty"`

	// 该目录下的文件条目数。
	ChildrenNum *int32 `json:"children_num,omitempty"`

	// 文件访问时间。
	AccessTime *int64 `json:"access_time,omitempty"`

	// 文件修改时间。
	ModificationTime *int64 `json:"modification_time,omitempty"`
}

func (o FileStatusV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FileStatusV2 struct{}"
	}

	return strings.Join([]string{"FileStatusV2", string(data)}, " ")
}
