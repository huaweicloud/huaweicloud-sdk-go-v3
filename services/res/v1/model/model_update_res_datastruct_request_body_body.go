package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateResDatastructRequestBodyBody This is a auto create Body Object
type UpdateResDatastructRequestBodyBody struct {

	// 物品特征信息。
	ItemAttrs []ItemAttrs `json:"item_attrs"`

	// 用户特征信息。
	UserAttrs []UserAttrs `json:"user_attrs"`

	Behaviors *BehaviorsConfig `json:"behaviors"`
}

func (o UpdateResDatastructRequestBodyBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResDatastructRequestBodyBody struct{}"
	}

	return strings.Join([]string{"UpdateResDatastructRequestBodyBody", string(data)}, " ")
}
