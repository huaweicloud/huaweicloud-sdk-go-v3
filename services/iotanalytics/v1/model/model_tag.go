package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Tag struct {

	// 标签名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 标签描述
	Description *string `json:"description,omitempty" xml:"description"`
}

func (o Tag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Tag struct{}"
	}

	return strings.Join([]string{"Tag", string(data)}, " ")
}
