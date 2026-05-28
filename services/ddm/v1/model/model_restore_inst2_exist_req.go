package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RestoreInst2ExistReq struct {
	Source *RestoreInstSource `json:"source"`

	Target *RestoreInstTarget `json:"target"`

	// 关联dn。
	DataNodeRelations []DataNodeRelation `json:"data_node_relations"`
}

func (o RestoreInst2ExistReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreInst2ExistReq struct{}"
	}

	return strings.Join([]string{"RestoreInst2ExistReq", string(data)}, " ")
}
