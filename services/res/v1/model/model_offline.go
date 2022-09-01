package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type Offline struct {

	// 用户数据url。
	UserUrl string `json:"user_url" xml:"user_url"`

	// 物品数据url。
	ItemUrl string `json:"item_url" xml:"item_url"`

	// 行为数据url。
	BehaviorUrl string `json:"behavior_url" xml:"behavior_url"`
}

func (o Offline) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Offline struct{}"
	}

	return strings.Join([]string{"Offline", string(data)}, " ")
}
