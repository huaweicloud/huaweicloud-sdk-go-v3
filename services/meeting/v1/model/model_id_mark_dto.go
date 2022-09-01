package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IdMarkDto struct {

	// 唯一标识。
	Id *string `json:"id,omitempty" xml:"id"`

	// id对应的回显描述，一般为名称等。
	Mark *string `json:"mark,omitempty" xml:"mark"`
}

func (o IdMarkDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IdMarkDto struct{}"
	}

	return strings.Join([]string{"IdMarkDto", string(data)}, " ")
}
