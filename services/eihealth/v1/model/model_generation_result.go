package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 分子生成任务的返回结果
type GenerationResult struct {

	// 任务名
	Name string `json:"name"`

	// 总生成轮数
	NumRounds int32 `json:"num_rounds"`

	// 期望条目数
	NumExpected int32 `json:"num_expected"`

	// 强约束数量
	NumStrongConstraints int32 `json:"num_strong_constraints"`

	// 弱约束数量
	NumWeakConstraints int32 `json:"num_weak_constraints"`

	// 分子ADMET属性名列表
	PropNames []string `json:"prop_names"`

	// 分子生成结果条目
	Result []GenerationResultItem `json:"result"`
}

func (o GenerationResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GenerationResult struct{}"
	}

	return strings.Join([]string{"GenerationResult", string(data)}, " ")
}
