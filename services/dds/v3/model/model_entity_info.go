package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EntityInfo struct {

	// 组ID或节点ID。
	EntityId string `json:"entity_id"`

	// 组名称或节点名称。
	EntityName string `json:"entity_name"`
}

func (o EntityInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EntityInfo struct{}"
	}

	return strings.Join([]string{"EntityInfo", string(data)}, " ")
}
