package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PersistObjectIdModifierDto struct {

	// 唯一标识。
	Id int64 `json:"id"`

	// 修改人。
	Modifier *string `json:"modifier,omitempty"`
}

func (o PersistObjectIdModifierDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PersistObjectIdModifierDto struct{}"
	}

	return strings.Join([]string{"PersistObjectIdModifierDto", string(data)}, " ")
}
