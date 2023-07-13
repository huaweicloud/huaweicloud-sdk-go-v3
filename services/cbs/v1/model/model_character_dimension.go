package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CharacterDimension 角色的尺寸
type CharacterDimension struct {

	// 角色的高度
	Height *string `json:"height,omitempty"`

	// 角色的宽度
	Width *string `json:"width,omitempty"`
}

func (o CharacterDimension) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CharacterDimension struct{}"
	}

	return strings.Join([]string{"CharacterDimension", string(data)}, " ")
}
