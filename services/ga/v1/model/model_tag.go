package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Tag 标签。
type Tag struct {

	// 标签的KEY。
	Key string `json:"key"`

	// 标签的VALUES。
	Values []string `json:"values"`
}

func (o Tag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Tag struct{}"
	}

	return strings.Join([]string{"Tag", string(data)}, " ")
}
