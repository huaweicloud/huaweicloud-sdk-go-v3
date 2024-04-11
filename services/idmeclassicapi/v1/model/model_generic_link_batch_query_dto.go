package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GenericLinkBatchQueryDto struct {

	// 是否返回源模型数据实例关联的最新版本目标模型数据实例。此参数仅对源/目标模型为M-V模型实体有效。 - true：返回源模型数据实例关联的最新版本的目标模型数据实例。 - false：返回源模型数据实例关联的所有版本的目标模型数据实例。默认为false。
	LatestOnly *bool `json:"latestOnly,omitempty"`

	// 角色对应的数据实例ID列表。
	ObjectIds *[]string `json:"objectIds,omitempty"`

	// 角色。 - TARGET：目标模型。 - SOURCE：源模型。
	Role *string `json:"role,omitempty"`
}

func (o GenericLinkBatchQueryDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GenericLinkBatchQueryDto struct{}"
	}

	return strings.Join([]string{"GenericLinkBatchQueryDto", string(data)}, " ")
}
