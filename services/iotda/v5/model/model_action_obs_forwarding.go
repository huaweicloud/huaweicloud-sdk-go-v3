package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 转发OBS服务消息结构
type ActionObsForwarding struct {
	// **参数说明**：OBS服务对应的region区域。

	RegionName string `json:"region_name"`
	// **参数说明**：OBS服务对应的projectId信息。

	ProjectId string `json:"project_id"`
	// **参数说明**：OBS服务对应的桶名称。

	BucketName string `json:"bucket_name"`
	// **参数说明**：OBS服务对应桶的区域。

	Location *string `json:"location,omitempty"`
}

func (o ActionObsForwarding) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionObsForwarding struct{}"
	}

	return strings.Join([]string{"ActionObsForwarding", string(data)}, " ")
}
