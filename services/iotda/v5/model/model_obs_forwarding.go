package model

import (
	"encoding/json"

	"strings"
)

// 服务配置信息
type ObsForwarding struct {
	// OBS服务对应的region区域
	RegionName string `json:"region_name"`
	// OBS服务对应的projectId信息
	ProjectId string `json:"project_id"`
	// OBS服务对应的桶名称
	BucketName string `json:"bucket_name"`
	// OBS服务对应桶的区域
	Location *string `json:"location,omitempty"`
}

func (o ObsForwarding) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ObsForwarding struct{}"
	}

	return strings.Join([]string{"ObsForwarding", string(data)}, " ")
}
