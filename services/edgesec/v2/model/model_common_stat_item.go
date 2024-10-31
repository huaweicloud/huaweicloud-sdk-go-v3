package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CommonStatItem 单个统计模型
type CommonStatItem struct {

	// 对应请求参数group_by的子类别。例如在Http攻击分布统计接口中，group_by为action时，key可为：log、block、captcha、js_challenge；在Http攻击Top接口中，group_by为url时，key可为请求的URL，例：/abc。
	Key string `json:"key"`

	// 攻击请求数
	Value int64 `json:"value"`
}

func (o CommonStatItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommonStatItem struct{}"
	}

	return strings.Join([]string{"CommonStatItem", string(data)}, " ")
}
