package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type CharacterConfig struct {

	// 形象id
	CharacterId string `json:"character_id"`

	Position *Position `json:"position,omitempty"`
}

func (o CharacterConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CharacterConfig struct{}"
	}

	return strings.Join([]string{"CharacterConfig", string(data)}, " ")
}
