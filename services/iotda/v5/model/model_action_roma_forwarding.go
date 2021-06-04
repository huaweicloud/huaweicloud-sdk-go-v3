package model

import (
	"encoding/json"

	"strings"
)

// 转发ROMA Connect服务消息结构
type ActionRomaForwarding struct {
	// **参数说明**：ROMA Connect服务对应的region区域。

	RegionName string `json:"region_name"`
	// **参数说明**：ROMA Connect服务对应的projectId信息。

	ProjectId string `json:"project_id"`
	// **参数说明**：ROMA Connect服务对应参数类型。

	RomaPushType *string `json:"roma_push_type,omitempty"`
}

func (o ActionRomaForwarding) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ActionRomaForwarding struct{}"
	}

	return strings.Join([]string{"ActionRomaForwarding", string(data)}, " ")
}
