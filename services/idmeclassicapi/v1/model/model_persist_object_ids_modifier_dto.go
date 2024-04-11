package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PersistObjectIdsModifierDto struct {

	// 数据实例ID列表。
	Ids []string `json:"ids"`

	// 修改人。
	Modifier *string `json:"modifier,omitempty"`
}

func (o PersistObjectIdsModifierDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PersistObjectIdsModifierDto struct{}"
	}

	return strings.Join([]string{"PersistObjectIdsModifierDto", string(data)}, " ")
}
