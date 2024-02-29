package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataEntityWithExtInfo 资产详情
type DataEntityWithExtInfo struct {
	Entity *Entity `json:"entity,omitempty"`

	// 关联资产信息，结构体Map<String, Entity>
	ReferredEntities *interface{} `json:"referred_entities,omitempty"`
}

func (o DataEntityWithExtInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataEntityWithExtInfo struct{}"
	}

	return strings.Join([]string{"DataEntityWithExtInfo", string(data)}, " ")
}
