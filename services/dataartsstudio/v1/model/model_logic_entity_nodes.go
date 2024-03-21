package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogicEntityNodes 业务资产目录树
type LogicEntityNodes struct {

	// 业务资产guid
	LogicEntityGuid *string `json:"logic_entity_guid,omitempty"`

	// 业务资产名称
	LogicEntityName *string `json:"logic_entity_name,omitempty"`
}

func (o LogicEntityNodes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogicEntityNodes struct{}"
	}

	return strings.Join([]string{"LogicEntityNodes", string(data)}, " ")
}
