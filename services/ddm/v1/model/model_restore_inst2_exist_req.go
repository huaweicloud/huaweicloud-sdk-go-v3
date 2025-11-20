package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RestoreInst2ExistReq struct {

	// 数据恢复源。
	Source *interface{} `json:"source,omitempty"`

	// 数据恢复目标。
	Target *interface{} `json:"target,omitempty"`

	// 关联dn。
	DataNodeRelations *[]DataNodeRelation `json:"data_node_relations,omitempty"`
}

func (o RestoreInst2ExistReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreInst2ExistReq struct{}"
	}

	return strings.Join([]string{"RestoreInst2ExistReq", string(data)}, " ")
}
