package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NameAndId 对象编号与名称
type NameAndId struct {

	// 对象编号
	Id *string `json:"id,omitempty"`

	// 对象名称
	Name *string `json:"name,omitempty"`
}

func (o NameAndId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NameAndId struct{}"
	}

	return strings.Join([]string{"NameAndId", string(data)}, " ")
}
