package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CharacterPosition struct {

	// 从左上为0点出发的横坐标
	X *int32 `json:"x,omitempty"`

	// 从左上为0点出发的纵坐标
	Y *string `json:"y,omitempty"`
}

func (o CharacterPosition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CharacterPosition struct{}"
	}

	return strings.Join([]string{"CharacterPosition", string(data)}, " ")
}
