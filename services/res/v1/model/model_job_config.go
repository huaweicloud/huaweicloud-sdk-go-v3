package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type JobConfig struct {
	NearlineRecallParam *NearLineRecallParam `json:"nearline_recall_param,omitempty" xml:"nearline_recall_param"`

	// 最大候选集个数（所有召回作业需提供此参数）。
	MaxRecommendedNum *int32 `json:"max_recommended_num,omitempty" xml:"max_recommended_num"`

	// 匹配特征对（属性匹配召回作业需要提供此参数）。
	MatchFeaturePairs *[]MatchFeaturePair `json:"match_feature_pairs,omitempty" xml:"match_feature_pairs"`

	Striping *Striping `json:"striping,omitempty" xml:"striping"`

	// 匹配类型（属性匹配召回作业需提供此参数）： - UI，基于用户推荐物品 - UU，基于用户推荐用户 - II，基于物品推荐物品 - IU，基于物品推荐用户
	MatchType *string `json:"match_type,omitempty" xml:"match_type"`

	MatrixFactorization *MatrixFactorization `json:"matrix_factorization,omitempty" xml:"matrix_factorization"`

	// 行为频率信息（历史行为记忆召回作业、历史行为过滤作业需提供此参数）。
	BehaviorFrequencys *[]BehaviorFrequency `json:"behavior_frequencys,omitempty" xml:"behavior_frequencys"`

	// 文件路径（人工配置候选集作业需要提供此参数）。
	FilePath *string `json:"file_path,omitempty" xml:"file_path"`

	UcbParam *UcbParam `json:"ucb_param,omitempty" xml:"ucb_param"`

	BehaviorGravity *BehaviorGravity `json:"behavior_gravity,omitempty" xml:"behavior_gravity"`

	Category *Category `json:"category,omitempty" xml:"category"`

	// 行为逻辑过滤（历史行为过滤作业需提供此参数）： - AND，同时满足则过滤 - OR， 满足一个则过滤
	BehaviorLogic *string `json:"behavior_logic,omitempty" xml:"behavior_logic"`

	FeaturesEngineering *EtlBasicParameter `json:"features_engineering,omitempty" xml:"features_engineering"`

	SampleParam *SampleParam `json:"sample_param,omitempty" xml:"sample_param"`

	DeepLearningParameters *DeepLearingParam `json:"deep_learning_parameters,omitempty" xml:"deep_learning_parameters"`

	AlgorithmSpecifyParameters *AlgorithmSpecifyParameters `json:"algorithm_specify_parameters,omitempty" xml:"algorithm_specify_parameters"`

	// 导入宽表（离线数据导入作业需要提供此参数）。
	LoadWidetable *bool `json:"load_widetable,omitempty" xml:"load_widetable"`

	// 导入画像（离线数据导入作业需要提供此参数）。
	LoadProfile *bool `json:"load_profile,omitempty" xml:"load_profile"`

	// 保留已有宽表（离线数据导入作业需要提供此参数）： - append，是 - new，否 - overwirte，覆盖
	SaveMode *string `json:"save_mode,omitempty" xml:"save_mode"`

	// 统计指标（效果评估作业需要提供此参数）。
	Indicators *[]Indicator `json:"indicators,omitempty" xml:"indicators"`

	// 离线排序作业名称（在线训练任务需要提供此参数）。
	OfflineRankJobName *string `json:"offline_rank_job_name,omitempty" xml:"offline_rank_job_name"`

	// 更新周期（在线训练任务需要提供此参数）。
	UpdateInterval *int32 `json:"update_interval,omitempty" xml:"update_interval"`

	Optimizer *Optimizer `json:"optimizer,omitempty" xml:"optimizer"`

	Flows *Flow `json:"flows,omitempty" xml:"flows"`
}

func (o JobConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobConfig struct{}"
	}

	return strings.Join([]string{"JobConfig", string(data)}, " ")
}
