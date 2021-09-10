package model

import (
	"encoding/json"

	"strings"
)

// 任务信息
type ClusterTask struct {
	// 任务描述

	Description *string `json:"description,omitempty"`
	// 任务id

	Id *string `json:"id,omitempty"`
	// 任务名称

	Name *string `json:"name,omitempty"`
}

func (o ClusterTask) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ClusterTask struct{}"
	}

	return strings.Join([]string{"ClusterTask", string(data)}, " ")
}
