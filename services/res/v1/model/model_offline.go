package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Offline
type Offline struct {

	// 用户数据url。
	UserUrl string `json:"user_url"`

	// 物品数据url。
	ItemUrl string `json:"item_url"`

	// 行为数据url。
	BehaviorUrl string `json:"behavior_url"`
}

func (o Offline) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Offline struct{}"
	}

	return strings.Join([]string{"Offline", string(data)}, " ")
}
