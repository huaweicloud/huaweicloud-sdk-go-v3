package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StorageMap struct {

	// **参数解释**： 存储大小。  **取值范围**： 不涉及。
	TotalStorage *string `json:"totalStorage,omitempty"`

	// **参数解释**： 文件数量。  **取值范围**： 不涉及。
	TotalCount *int32 `json:"totalCount,omitempty"`
}

func (o StorageMap) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StorageMap struct{}"
	}

	return strings.Join([]string{"StorageMap", string(data)}, " ")
}
