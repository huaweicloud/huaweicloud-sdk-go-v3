package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AiguardAtlasUpdateBody struct {

	// 图库名称
	Name *string `json:"name,omitempty"`

	// 图库描述
	Description *string `json:"description,omitempty"`
}

func (o AiguardAtlasUpdateBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiguardAtlasUpdateBody struct{}"
	}

	return strings.Join([]string{"AiguardAtlasUpdateBody", string(data)}, " ")
}
