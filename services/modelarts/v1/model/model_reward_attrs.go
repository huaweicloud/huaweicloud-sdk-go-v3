package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RewardAttrs 搜索指标列表。
type RewardAttrs struct {

	// 指标名称。
	Name *string `json:"name,omitempty"`

	// 搜索方向。 - max指定时表示指标值越大越好； - min指定时表示指标值越小越好。
	Mode *string `json:"mode,omitempty"`

	// 指标正则表达式。
	Regex *string `json:"regex,omitempty"`
}

func (o RewardAttrs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RewardAttrs struct{}"
	}

	return strings.Join([]string{"RewardAttrs", string(data)}, " ")
}
