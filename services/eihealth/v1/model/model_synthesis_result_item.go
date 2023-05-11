package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 分子合成路径规划结果条目
type SynthesisResultItem struct {

	// 分子合成规划，列表内是reactions id
	Route []string `json:"route"`

	// 当前分子合成路径的得分
	Score float32 `json:"score"`
}

func (o SynthesisResultItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SynthesisResultItem struct{}"
	}

	return strings.Join([]string{"SynthesisResultItem", string(data)}, " ")
}
