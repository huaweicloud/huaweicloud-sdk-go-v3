package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Tag struct {
	// 是否与分支重名

	IsDoubleName *bool `json:"is_double_name,omitempty"`
	// 标签名

	Name *string `json:"name,omitempty"`
}

func (o Tag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Tag struct{}"
	}

	return strings.Join([]string{"Tag", string(data)}, " ")
}
