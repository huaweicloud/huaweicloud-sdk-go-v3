package model

import (
	"encoding/json"

	"strings"
)

type ShowCeshierarchyRespQueues1 struct {
	// topic名称。

	Name *string `json:"name,omitempty"`
	// 分区信息。

	Partitions *[]ShowCeshierarchyRespPartitions `json:"partitions,omitempty"`
}

func (o ShowCeshierarchyRespQueues1) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowCeshierarchyRespQueues1 struct{}"
	}

	return strings.Join([]string{"ShowCeshierarchyRespQueues1", string(data)}, " ")
}
