package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StorageId **参数解释**：存储资源ID，可通过[查询存储资源列表](ListV2Storages.xml)获取存储资源ID。 **约束限制**：不涉及。 **取值范围**：长度为1~36的英文字符、数字和中划线的组合。 **默认取值**：不涉及。
type StorageId struct {
}

func (o StorageId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StorageId struct{}"
	}

	return strings.Join([]string{"StorageId", string(data)}, " ")
}
