package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFtDetailResponse Response Object
type ShowFtDetailResponse struct {

	// 项目id。
	ProjectId *string `json:"project_id,omitempty"`

	// 训练任务id。
	TaskId *string `json:"task_id,omitempty"`

	// 训练任务名称。
	TaskName *string `json:"task_name,omitempty"`

	// 训练任务描述信息。
	TaskDesc *string `json:"task_desc,omitempty"`

	Metadata *JobMetadataResponse `json:"metadata,omitempty"`

	Spec *SpecResponse `json:"spec,omitempty"`

	// 模型id。
	ModelAssetId *string `json:"model_asset_id,omitempty"`

	// **参数解释：** 模型类型，取值为TextGeneration|ImageUnderstanding，依次为：文本生成、图像理解。 **约束限制：** 不涉及 **取值范围：** TextGeneration|ImageUnderstanding **默认取值：** 不涉及
	ModelType *string `json:"model_type,omitempty"`

	// 模型来源
	ModelSource *string `json:"model_source,omitempty"`

	// **参数解释：** 训练类型，支持SFT（全量微调）、PRETRAIN（预训练）、LORA（lora微调）、DPO（dpo强化学习）、RFT（rft强化学习）。 **约束限制：** 不涉及 **取值范围：** SFT（全量微调）、PRETRAIN（预训练）、LORA（lora微调）、DPO（dpo强化学习）、RFT（rft强化学习） 默认取值： SFT
	TrainType *string `json:"train_type,omitempty"`

	// 断点续训相关配置。
	CheckpointConfig *string `json:"checkpoint_config,omitempty"`

	// 训练任参数信息。
	TaskParameters *string `json:"task_parameters,omitempty"`

	// 创建时间。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 训练任务更新时间，当修改、或者训练任务状态发生变化时进行更新。
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 训练任务进度。
	TrainProcess *float64 `json:"train_process,omitempty"`

	// 该训练任务数据集相关的配置。
	DatasetsConfig *[]DatasetConfig `json:"datasets_config,omitempty"`

	Status *Status `json:"status,omitempty"`

	// 自动发布配置信息
	AutoPublishConfig *string `json:"auto_publish_config,omitempty"`

	// 模型资产名
	AssetCode *string `json:"asset_code,omitempty"`

	// 资产名称
	AssetName *string `json:"asset_name,omitempty"`

	// 模型资产描述信息
	AssetDesc *string `json:"asset_desc,omitempty"`

	// 模型系列
	AssetSeries *string `json:"asset_series,omitempty"`

	// 资产版本
	AssetVersion *string `json:"asset_version,omitempty"`

	// 资产类型
	AssetType *string `json:"asset_type,omitempty"`

	// 资产来源
	AssetSource *string `json:"asset_source,omitempty"`

	// 资产组id
	AssetGroupId *string `json:"asset_group_id,omitempty"`

	// 资产子类型
	SubAssetType *string `json:"sub_asset_type,omitempty"`

	// 资产类别
	Category *string `json:"category,omitempty"`

	// 资产API版本
	ApiVersion *string `json:"api_version,omitempty"`

	// 根资产ID
	RootAssetId *string `json:"root_asset_id,omitempty"`

	// 训练任务耗时
	TrainCostTime *int64 `json:"train_cost_time,omitempty"`

	// 任务所属工作空间名称
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 用户id
	UserId *string `json:"user_id,omitempty"`

	// 用户名称
	UserName *string `json:"user_name,omitempty"`

	// 资源池类型
	PoolType *string `json:"pool_type,omitempty"`

	// 资源池ID
	PoolId *string `json:"pool_id,omitempty"`

	// 使用的资源池实例数
	PoolNodeCount *string `json:"pool_node_count,omitempty"`

	// 使用的资源池卡数
	FlavorId *string `json:"flavor_id,omitempty"`

	// 优先级
	Priority *int32 `json:"priority,omitempty"`

	// 训练预估时长
	TrainingInfo *string `json:"training_info,omitempty"`

	// **参数解释**：训练产物输出路径，如\"obs://yyy/test/\"。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	TrainOutputPath *string `json:"train_output_path,omitempty"`

	// 训练模型类型
	AssetCapabilities *[]string `json:"asset_capabilities,omitempty"`

	ContinueTask   *ContinueTask `json:"continue_task,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowFtDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFtDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowFtDetailResponse", string(data)}, " ")
}
