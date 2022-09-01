package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type Nearline struct {
	ItemTopic *Topic `json:"item_topic,omitempty" xml:"item_topic"`

	UserTopic *Topic `json:"user_topic,omitempty" xml:"user_topic"`

	BehaviorTopic *Topic `json:"behavior_topic,omitempty" xml:"behavior_topic"`
}

func (o Nearline) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Nearline struct{}"
	}

	return strings.Join([]string{"Nearline", string(data)}, " ")
}
