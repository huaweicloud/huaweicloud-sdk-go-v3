package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionModelVersionReviseAndUpdateDtoVersionModel struct {

	// 创建人。
	Creator *string `json:"creator,omitempty"`

	// 关系实体名称集合，与workCopyType的值CUSTOM配合使用。
	CustomLinkSet *[]string `json:"customLinkSet,omitempty"`

	Data *VersionModel `json:"data"`

	// 主对象ID。
	MasterId string `json:"masterId"`

	// 更新者。
	Modifier *string `json:"modifier,omitempty"`

	// 关系的复制类型。 - BOTH：若存在关系实例引用此数据实例作为源端实例或目标端实例，修订且更新后的数据实例将继承这些关系实例。 - SOURCE：若存在关系实例引用此数据实例作为源端实例，修订且更新后的数据实例将继承这些关系实例。 - TARGET：若存在关系实例引用此数据实例作为目标端实例，修订且更新后的数据实例将继承这些关系实例。 - NONE：修订且更新后的数据实例将不继承任何关系实例。 - CUSTOM：若指定的关系实体集合对应的关系实例引用此数据实例作为源端实例或目标端实例，修订且更新后的数据实例将继承这些关系实例。
	WorkCopyType *string `json:"workCopyType,omitempty"`

	// 是否已检出。 - true：已检出。 - false：未检出。
	WorkingCopy *bool `json:"workingCopy,omitempty"`
}

func (o VersionModelVersionReviseAndUpdateDtoVersionModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionModelVersionReviseAndUpdateDtoVersionModel struct{}"
	}

	return strings.Join([]string{"VersionModelVersionReviseAndUpdateDtoVersionModel", string(data)}, " ")
}
