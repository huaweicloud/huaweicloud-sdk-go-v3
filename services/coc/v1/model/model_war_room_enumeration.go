package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WarRoomEnumeration 枚举类型信息
type WarRoomEnumeration struct {

	// 枚举值id
	Id *string `json:"id,omitempty"`

	// 枚举值中文名
	NameZh *string `json:"name_zh,omitempty"`

	// 枚举值英文名
	NameEn *string `json:"name_en,omitempty"`

	// 枚举类型
	Type *string `json:"type,omitempty"`
}

func (o WarRoomEnumeration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WarRoomEnumeration struct{}"
	}

	return strings.Join([]string{"WarRoomEnumeration", string(data)}, " ")
}
