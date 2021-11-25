package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 命名实体。
type NamedEntity struct {
	// 实体文本。

	Word string `json:"word"`
	// 实体类型，枚举类型。 中文支持人名“nr”，地名“ns”，机构名“nt”，时间“t”。 英文支持人名“per”，地名“loc”，机构名“org”，时间“t”。 西班牙文支持人名“per”，地名“loc”，机构名“org”，时间“t”。

	Tag string `json:"tag"`
	// 实体文本在待分析文本中的起始位置。

	Offset int32 `json:"offset"`
	// 实体文本长度。

	Len int32 `json:"len"`
}

func (o NamedEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NamedEntity struct{}"
	}

	return strings.Join([]string{"NamedEntity", string(data)}, " ")
}
