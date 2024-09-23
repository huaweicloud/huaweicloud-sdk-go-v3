package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EntityDto 包含有关根、组织单位或账号的信息。
type EntityDto struct {

	// 实体的名称。
	Name string `json:"name"`

	// 实体的唯一标识符（ID）。
	Id string `json:"id"`

	// 实体的类型。account：账号；organizational_unit：组织单元；root：根。
	Type string `json:"type"`
}

func (o EntityDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EntityDto struct{}"
	}

	return strings.Join([]string{"EntityDto", string(data)}, " ")
}
