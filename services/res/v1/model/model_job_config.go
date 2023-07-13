package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type JobConfig struct {
	NearlineRecallParam *NearLineRecallParam `json:"nearline_recall_param,omitempty"`

	// 最大候选集个数（所有召回作业需提供此参数）。
	MaxRecommendedNum *int32 `json:"max_recommended_num,omitempty"`

	// 匹配特征对（属性匹配召回作业需要提供此参数）。
	MatchFeaturePairs *[]MatchFeaturePair `json:"match_feature_pairs,omitempty"`

	Striping *Striping `json:"striping,omitempty"`

	// 匹配类型（属性匹配召回作业需提供此参数）： - UI，基于用户推荐物品 - UU，基于用户推荐用户 - II，基于物品推荐物品 - IU，基于物品推荐用户
	MatchType *string `json:"match_type,omitempty"`

	MatrixFactorization *MatrixFactorization `json:"matrix_factorization,omitempty"`

	// 行为频率信息（历史行为记忆召回作业、历史行为过滤作业需提供此参数）。
	BehaviorFrequencys *[]BehaviorFrequency `json:"behavior_frequencys,omitempty"`

	// 文件路径（人工配置候选集作业需要提供此参数）。
	FilePath *string `json:"file_path,omitempty"`

	UcbParam *UcbParam `json:"ucb_param,omitempty"`

	BehaviorGravity *BehaviorGravity `json:"behavior_gravity,omitempty"`

	Category *Category `json:"category,omitempty"`

	// 行为逻辑过滤（历史行为过滤作业需提供此参数）： - AND，同时满足则过滤 - OR， 满足一个则过滤
	BehaviorLogic *string `json:"behavior_logic,omitempty"`

	FeaturesEngineering *EtlBasicParameter `json:"features_engineering,omitempty"`

	SampleParam *SampleParam `json:"sample_param,omitempty"`

	DeepLearningParameters *DeepLearingParam `json:"deep_learning_parameters,omitempty"`

	AlgorithmSpecifyParameters *AlgorithmSpecifyParameters `json:"algorithm_specify_parameters,omitempty"`

	// 导入宽表（离线数据导入作业需要提供此参数）。
	LoadWidetable *bool `json:"load_widetable,omitempty"`

	// 导入画像（离线数据导入作业需要提供此参数）。
	LoadProfile *bool `json:"load_profile,omitempty"`

	// 保留已有宽表（离线数据导入作业需要提供此参数）： - append，是 - new，否 - overwirte，覆盖
	SaveMode *string `json:"save_mode,omitempty"`

	// 统计指标（效果评估作业需要提供此参数）。
	Indicators *[]Indicator `json:"indicators,omitempty"`

	// 离线排序作业名称（在线训练任务需要提供此参数）。
	OfflineRankJobName *string `json:"offline_rank_job_name,omitempty"`

	// 更新周期（在线训练任务需要提供此参数）。
	UpdateInterval *int32 `json:"update_interval,omitempty"`

	Optimizer *Optimizer `json:"optimizer,omitempty"`

	Flows *Flow `json:"flows,omitempty"`
}

func (o JobConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobConfig struct{}"
	}

	return strings.Join([]string{"JobConfig", string(data)}, " ")
}
