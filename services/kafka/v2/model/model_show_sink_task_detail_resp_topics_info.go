package model

import (
	"encoding/json"

	"strings"
)

type ShowSinkTaskDetailRespTopicsInfo struct {
	// topic名称。

	Topic *string `json:"topic,omitempty"`
	// 分区列表。

	Partitions *[]ShowSinkTaskDetailRespPartitions `json:"partitions,omitempty"`
}

func (o ShowSinkTaskDetailRespTopicsInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowSinkTaskDetailRespTopicsInfo struct{}"
	}

	return strings.Join([]string{"ShowSinkTaskDetailRespTopicsInfo", string(data)}, " ")
}
