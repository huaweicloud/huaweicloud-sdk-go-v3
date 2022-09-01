package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MigrateNodesTask struct {

	// API版本，固定值“v3”。
	ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion"`

	// API类型，固定值“MigrateNodesTask”。
	Kind *string `json:"kind,omitempty" xml:"kind"`

	Spec *MigrateNodesSpec `json:"spec" xml:"spec"`

	Status *TaskStatus `json:"status,omitempty" xml:"status"`
}

func (o MigrateNodesTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateNodesTask struct{}"
	}

	return strings.Join([]string{"MigrateNodesTask", string(data)}, " ")
}
