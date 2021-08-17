package model

import (
	"encoding/json"

	"strings"
)

// 推荐的目的端服务器配置
type InitTargetServer struct {
	// 推荐的目的端服务器的磁盘信息

	Disks []DiskIntargetServer `json:"disks"`
}

func (o InitTargetServer) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "InitTargetServer struct{}"
	}

	return strings.Join([]string{"InitTargetServer", string(data)}, " ")
}
