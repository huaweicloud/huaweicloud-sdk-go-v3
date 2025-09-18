package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BotMDefenseStrategy struct {
	Low *BotMDefenseLevel `json:"low,omitempty"`

	Medium *BotMDefenseLevel `json:"medium,omitempty"`

	High *BotMDefenseLevel `json:"high,omitempty"`
}

func (o BotMDefenseStrategy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BotMDefenseStrategy struct{}"
	}

	return strings.Join([]string{"BotMDefenseStrategy", string(data)}, " ")
}
