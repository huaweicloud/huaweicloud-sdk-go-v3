package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AnimationItem struct {

	// 动作资产ID。
	AnimationAssetId *string `json:"animation_asset_id,omitempty"`

	// 时间戳，相对时间戳。  单位秒。  保留3位小数。
	Timestamp *float32 `json:"timestamp,omitempty"`
}

func (o AnimationItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AnimationItem struct{}"
	}

	return strings.Join([]string{"AnimationItem", string(data)}, " ")
}
