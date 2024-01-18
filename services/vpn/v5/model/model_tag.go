package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Tag struct {
	Key string `json:"key"`

	Values []string `json:"values"`
}

func (o Tag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Tag struct{}"
	}

	return strings.Join([]string{"Tag", string(data)}, " ")
}
