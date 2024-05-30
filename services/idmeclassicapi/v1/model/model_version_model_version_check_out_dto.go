package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionModelVersionCheckOutDto struct {

	// 创建人。
	Creator *string `json:"creator,omitempty"`

	// 关系实体名称集合，与workCopyType的值CUSTOM配合使用。
	CustomLinkSet *[]string `json:"customLinkSet,omitempty"`

	// 主对象ID。
	MasterId string `json:"masterId"`

	// 更新者。
	Modifier *string `json:"modifier,omitempty"`

	// 关系的复制类型。 - BOTH：若存在关系实例引用此数据实例作为源端实例或目标端实例，检出后的数据实例将继承这些关系实例。 - SOURCE：若存在关系实例引用此数据实例作为源端实例，检出后的数据实例将继承这些关系实例。 - TARGET：若存在关系实例引用此数据实例作为目标端实例，检出后的数据实例将继承这些关系实例。 - NONE：检出后的数据实例将不继承任何关系实例。 - CUSTOM：若指定的关系实体集合对应的关系实例引用此数据实例作为源端实例或目标端实例，检出后的数据实例将继承这些关系实例。
	WorkCopyType *string `json:"workCopyType,omitempty"`
}

func (o VersionModelVersionCheckOutDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionModelVersionCheckOutDto struct{}"
	}

	return strings.Join([]string{"VersionModelVersionCheckOutDto", string(data)}, " ")
}
