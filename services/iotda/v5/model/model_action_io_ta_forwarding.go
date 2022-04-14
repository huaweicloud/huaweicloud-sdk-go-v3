package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 转发IoTA服务消息结构
type ActionIoTaForwarding struct {
	// **参数说明**：IoTA服务对应的region区域。

	RegionName string `json:"region_name"`
	// **参数说明**：IoTA服务对应的projectId信息。

	ProjectId string `json:"project_id"`
}

func (o ActionIoTaForwarding) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionIoTaForwarding struct{}"
	}

	return strings.Join([]string{"ActionIoTaForwarding", string(data)}, " ")
}
