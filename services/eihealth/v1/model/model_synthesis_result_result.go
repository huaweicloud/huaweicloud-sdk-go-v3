package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SynthesisResultResult 合成路径规划结果字典
type SynthesisResultResult struct {

	// 分子合成规划中的分子
	Molecules []SynthesisResultResultMolecules `json:"molecules"`

	// 分子合成规划中的反应列表
	Reactions []SynthesisResultResultReactions `json:"reactions"`

	// 分子合成规划的具体信息
	SynthesisRoutes []SynthesisResultItem `json:"synthesis_routes"`
}

func (o SynthesisResultResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SynthesisResultResult struct{}"
	}

	return strings.Join([]string{"SynthesisResultResult", string(data)}, " ")
}
