package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IdMarkDto struct {
	// 唯一标识。

	Id *string `json:"id,omitempty"`
	// id对应的回显描述，一般为名称等。

	Mark *string `json:"mark,omitempty"`
}

func (o IdMarkDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IdMarkDto struct{}"
	}

	return strings.Join([]string{"IdMarkDto", string(data)}, " ")
}
