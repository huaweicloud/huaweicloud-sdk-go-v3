package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LinkedEntity 链接的实体 类
type LinkedEntity struct {

	// 实体指称
	Mention string `json:"mention"`

	// 偏移量
	Offset int32 `json:"offset"`

	// 实体名称
	EntityTitle string `json:"entity_title"`
}

func (o LinkedEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LinkedEntity struct{}"
	}

	return strings.Join([]string{"LinkedEntity", string(data)}, " ")
}
