package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 标签详情。
type ProjectTag struct {

	// 键。
	Key string `json:"key"`

	// 值。
	Values []string `json:"values"`
}

func (o ProjectTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectTag struct{}"
	}

	return strings.Join([]string{"ProjectTag", string(data)}, " ")
}
