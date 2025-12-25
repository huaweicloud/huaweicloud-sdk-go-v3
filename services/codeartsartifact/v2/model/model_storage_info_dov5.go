package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StorageInfoDov5 struct {

	// **参数解释**: 使用空间（字节）。 **取值范围**: 不涉及。
	UsedSpaceLength *int64 `json:"used_space_length,omitempty"`

	// **参数解释**: 使用空间（带单位）。 **取值范围**: 不涉及。
	UsedSpace *string `json:"used_space,omitempty"`

	// **参数解释**: 文件数量。 **取值范围**: 不涉及。
	FilesCount *int64 `json:"files_count,omitempty"`
}

func (o StorageInfoDov5) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StorageInfoDov5 struct{}"
	}

	return strings.Join([]string{"StorageInfoDov5", string(data)}, " ")
}
