package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GenericLinkTypeModifierDto struct {

	// 是否返回源模型数据实例关联的最新版本目标模型数据实例。此参数仅对源/目标模型为M-V模型实体有效。 - true：返回源模型数据实例关联的最新版本的目标模型数据实例。 - false：返回源模型数据实例关联的所有版本的目标模型数据实例。默认为false。
	LatestOnly *bool `json:"latestOnly,omitempty"`

	// 修改人。
	Modifier *string `json:"modifier,omitempty"`

	// 源模型数据实例的ID。
	SourceId *string `json:"sourceId,omitempty"`

	// 目标模型的英文名称。
	TargetType *string `json:"targetType,omitempty"`
}

func (o GenericLinkTypeModifierDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GenericLinkTypeModifierDto struct{}"
	}

	return strings.Join([]string{"GenericLinkTypeModifierDto", string(data)}, " ")
}
