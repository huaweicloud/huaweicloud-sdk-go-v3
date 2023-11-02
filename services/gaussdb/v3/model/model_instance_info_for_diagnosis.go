package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceInfoForDiagnosis struct {

	// 实例ID。
	InstanceId *string `json:"instance_id,omitempty"`

	// 主节点ID。
	MasterNodeId *string `json:"master_node_id,omitempty"`
}

func (o InstanceInfoForDiagnosis) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceInfoForDiagnosis struct{}"
	}

	return strings.Join([]string{"InstanceInfoForDiagnosis", string(data)}, " ")
}
