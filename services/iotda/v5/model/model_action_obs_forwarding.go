package model

import (
	"encoding/json"

	"strings"
)

// 转发OBS服务消息结构
type ActionObsForwarding struct {
	// OBS服务对应的region区域

	RegionName string `json:"region_name"`
	// OBS服务对应的projectId信息

	ProjectId string `json:"project_id"`
	// OBS服务对应的桶名称

	BucketName string `json:"bucket_name"`
	// OBS服务对应桶的区域

	Location *string `json:"location,omitempty"`
}

func (o ActionObsForwarding) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ActionObsForwarding struct{}"
	}

	return strings.Join([]string{"ActionObsForwarding", string(data)}, " ")
}
