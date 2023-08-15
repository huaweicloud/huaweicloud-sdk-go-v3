package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AvailableTag 可用区标签
type AvailableTag struct {

	// 可用区计费模式，分为专属dedicated和共享shard
	Mode *string `json:"mode,omitempty"`

	// az的别名
	Alias *string `json:"alias,omitempty"`

	// 所属group。默认为”center”
	PublicBorderGroup *string `json:"public_border_group,omitempty"`
}

func (o AvailableTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AvailableTag struct{}"
	}

	return strings.Join([]string{"AvailableTag", string(data)}, " ")
}
