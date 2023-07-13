package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EntityWithExtInfo 资产详情
type EntityWithExtInfo struct {
	Entity *AtlasAssetEntity `json:"entity"`

	// 关联资产map Map<String, AtlasAssetEntity>
	ReferredEntities *interface{} `json:"referred_entities,omitempty"`
}

func (o EntityWithExtInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EntityWithExtInfo struct{}"
	}

	return strings.Join([]string{"EntityWithExtInfo", string(data)}, " ")
}
