package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Nearline
type Nearline struct {
	ItemTopic *Topic `json:"item_topic,omitempty"`

	UserTopic *Topic `json:"user_topic,omitempty"`

	BehaviorTopic *Topic `json:"behavior_topic,omitempty"`
}

func (o Nearline) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Nearline struct{}"
	}

	return strings.Join([]string{"Nearline", string(data)}, " ")
}
