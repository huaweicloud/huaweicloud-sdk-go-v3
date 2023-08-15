package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateResJobRequestBody This is a auto create Body Object
type UpdateResJobRequestBody struct {

	// 类别： - RECALL，召回作业 - DATASOURCE，数据源作业 - FILTER，过滤作业 - SORTING，排序作业 - EVALUATE，效果评估作业
	Category string `json:"category"`

	// 描述
	Description *string `json:"description,omitempty"`

	JobConfig *JobConfig `json:"job_config"`

	// 作业名称
	JobName string `json:"job_name"`

	// 作业类型： - WriteUserProfile，用户画像实时导入作业 - WriteItemProfile，物品画像实时导入作业 - UniversalProcess，行为数据实时导入作业 - NearlineRecall，近线召回作业 - EncodeProfile，近线特征工程作业 - AttributeMatch，属性匹配召回作业 - AlsCF，交替最小二乘协同过滤作业 - BhvHistory，历史行为记忆召回作业 - ItemCf，物品协同过滤召回作业 - MenEdit，人工配置候选集作业 - Ucb，UCB召回作业 - UserCf，用户协同过滤召回作业 - WeightBehavior，综合行为热度召回作业 - Filter，历史行为过滤作业 - AutoPreRank，智能ETL参数生成作业 - ETL，离线特征工程作业 - LR，LR作业 - DEEPFM，DEEPFM作业 - AutoGroup，AutoGroup作业 - StreamRank，在线训练作业 - DataStruct，识别数据结构作业 - DataExploration，数据探索作业 - DataImport，离线数据导入作业 - Evaluate，效果评估作业
	JobType string `json:"job_type"`

	// 调度参数
	Schedule *string `json:"schedule,omitempty"`
}

func (o UpdateResJobRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResJobRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateResJobRequestBody", string(data)}, " ")
}
