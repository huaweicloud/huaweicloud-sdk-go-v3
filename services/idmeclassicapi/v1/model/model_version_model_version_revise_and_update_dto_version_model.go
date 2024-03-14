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

	// 关系的复制类型。 - BOTH：复制当前M-V模型作为源端与目标端的关系。 - CUSTOM：自定义复制当前M-V模型的关系。 - NONE：不复制当前M-V模型的关系。 - SOURCE：仅复制当前M-V模型作为源端的关系。 - TARGET：仅复制当前M-V模型作为目标端的关系。
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
