package model

import (
	"encoding/json"

	"strings"
)

// 节点的磁盘信息，请参见volume参数说明。
type ClusterDetailInstanceVolume struct {
	// 节点的磁盘类型，只支持本地磁盘

	Type *string `json:"type,omitempty"`
	// 节点磁盘大小，单位G。

	Size *int64 `json:"size,omitempty"`
}

func (o ClusterDetailInstanceVolume) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ClusterDetailInstanceVolume struct{}"
	}

	return strings.Join([]string{"ClusterDetailInstanceVolume", string(data)}, " ")
}
