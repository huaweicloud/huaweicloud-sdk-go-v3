package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAssetRequest Request Object
type ListAssetRequest struct {

	// **参数解释**：排序方向。 **约束限制**：不涉及。 **取值范围**： 枚举值： - ASC: 升序 - DESC: 降序。 **默认取值**：DESC。
	SortDir *string `json:"sort_dir,omitempty"`

	// **参数解释**：排序字段。 **约束限制**：不涉及。 **取值范围**：枚举值： - created_at: 创建时间 - updated_at: 更新时间。 **默认取值**：created_at。
	SortKey *string `json:"sort_key,omitempty"`

	// **参数解释**：限制量，单次查询总量，必须由数字组成。 **约束限制**：不涉及。 **取值范围**：1-1000。 **默认取值**：100。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：偏移量，查询起始偏移，必须由数字组成。 **约束限制**：不涉及。 **取值范围**0-100000000。 **默认取值**：0。
	Offset *int32 `json:"offset,omitempty"`

	// 工作空间ID。[获取方法请参见[查询工作空间列表](ListWorkspace.xml)。](tag:hc)未创建工作空间时默认值为“0”，存在创建并使用的工作空间，以实际取值为准。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// **参数解释**：资产名称,支持英文字母，数字，中划线，下划线，点，空格，中文汉字。 **约束限制**：1-128位字符。 **取值范围**：不涉及。 **默认取值**：不涉及。
	AssetName *string `json:"asset_name,omitempty"`

	// **参数解释**: 资产类型。 **约束限制**: 枚举值: - MODEL: 模型资产 - TEMPLATE: 模板资产 - DATASET: 数据集资产 - WORKFLOW: 工作流资产 - IMAGE: 镜像资产 - ALGORITHM: 算法资产 - DATASET_TEMPLATE: 数据集模板资产 - DATA_OPERATOR: 算子镜像资产 **默认取值**: 不涉及。
	AssetType *[]string `json:"asset_type,omitempty"`

	// **参数解释**：资产来源。 **约束限制**：不涉及。 **取值范围**：枚举值: - Preset: 预置资产 - Publish: 用户发布资产 - Import: 用户导入资产 **默认取值**：不涉及。
	AssetSource *string `json:"asset_source,omitempty"`

	// **参数解释**：用户资产来源。 **约束限制**：1-60位字符，英文字母，数字，下划线、中划线。 **取值范围**：preset_train：精调，custom_train：自定义训练，AgenticRL：强化学习。 **默认取值**：不涉及。
	AssetSourceType *string `json:"asset_source_type,omitempty"`

	// **参数解释**：资产编码。 **约束限制**：1-60字母，数字，中划线，下划线，点。 **取值范围**：不涉及。 **默认取值**：不涉及。
	AssetCode *string `json:"asset_code,omitempty"`

	// **参数解释**：资产版本。 **约束限制**：数字或格式为x.y.z，x,y,z为数字。 **取值范围**：不涉及。 **默认取值**：不涉及。
	AssetVersion *string `json:"asset_version,omitempty"`

	// **参数解释**：资产归属的资产组id，带有该字段将在同一个资产库中所有版本进行条件筛选，否则会筛选所有资产库中的主版本。 **约束限制**：1-36字母数字，下划线，中划线。 **取值范围**：不涉及。 **默认取值**：不涉及。
	AssetGroupId *string `json:"asset_group_id,omitempty"`

	// **参数解释**：是否过滤所有版本。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：false。
	FilterAllVersion *bool `json:"filter_all_version,omitempty"`

	// **参数解释**：资产支持的操作列表，例如SFT：监督微调，LORA微调:LORA，在线部署：ONLINE-DEPLOY，强化学习：AgenticRL。 **约束限制**：每个元素1-30位字符，英文字母，数字，下划线，中划线；最多10个元素。 **取值范围**：不涉及。 **默认取值**：不涉及。
	SupportActions *[]string `json:"support_actions,omitempty"`

	// **参数解释**：检索资产时的标签，根据不同资产类型，条件可能不同。 - 需要匹配一个key对应一个value时格式为key=value - 需要匹配一个key对应多个value中的任意一个时格式为key=value1|value2|value3 - 需要匹配多个条件同时满足时key1=value1,key2=value2 - 如果需要实现模糊匹配格式为key1%=value2（右匹配）key2=%value2（左匹配） **约束限制**：不涉及。
	Tag *[]string `json:"tag,omitempty"`

	// **参数解释**：过滤我创建的资产，boolean类型，取值true：查询我创建的资产，取值false：查询所有资产。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	IsMine *bool `json:"is_mine,omitempty"`

	// **参数解释**：列表查询是否包含action。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	WithActions *bool `json:"with_actions,omitempty"`

	// **参数解释**：资产品牌，例如DeepSeek,Qwen,Wan,GPT,Llama,GLM,ERNIE,Mistral,Kimi,Other。 **约束限制**：不涉及。 **取值范围**：0-100位。 **默认取值**：不涉及。
	AssetSeries *string `json:"asset_series,omitempty"`

	// **参数解释**：资产能力，为自由文本字符串列表(非严格枚举)。推荐使用以下约定值：例如TextGeneration,ImageUnderstanding,ImageGeneration,VideoGeneration,AgenticRL,Other。 **约束限制**：每个元素为1-40位字符；最多10个元素。
	AssetCapabilities *[]string `json:"asset_capabilities,omitempty"`

	// **参数解释**：资产场景，如DeepThinking。 **约束限制**：不涉及。
	AssetScenes *[]string `json:"asset_scenes,omitempty"`

	// **参数解释**：根资产id。 **约束限制**：1-36个字符，英文字母，数字，中划线，下划线。 **取值范围**：不涉及。 **默认取值**：不涉及。
	RootAssetId *string `json:"root_asset_id,omitempty"`

	// **参数解释**：父资产id。 **约束限制**：1-36个字符，英文字母，数字，中划线，下划线。 **取值范围**：不涉及。 **默认取值**：不涉及。
	ParentAssetId *string `json:"parent_asset_id,omitempty"`

	// **参数解释**：是否返回软删除的资产。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：false。
	IsDeleted *bool `json:"is_deleted,omitempty"`

	// **参数解释**：资产id。 **约束限制**：0-100个。
	AssetIds *[]string `json:"asset_ids,omitempty"`

	// **参数解释**：tag组合方式 **约束限制**：不涉及。 **取值范围**：and：与，or：或。 **默认取值**：and。
	TagMode *string `json:"tag_mode,omitempty"`

	// **参数解释**：数据集用途，例如 GENERAL：通用数据集、EVAL：评测集。 **约束限制**：1-40位字符，英文字母，数字，下划线，中划线。 **取值范围**：不涉及。 **默认取值**：不涉及。
	DatasetUsage *string `json:"dataset_usage,omitempty"`

	// **参数解释**：资产状态。 **约束限制**：不涉及。 **取值范围**：枚举值: - online: 上线 - offline: 下线 - waiting: 等待中 **默认取值**：不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释**：数据集存储格式，例如JSONL，TXT，MOBI，EPUB，DOCX，PDF。 **约束限制**：不涉及。
	DatasetFileFormats *[]string `json:"dataset_file_formats,omitempty"`

	// **参数解释**：数据集模态划分，如音频，图像，视频。 **约束限制**：1-40位字符，英文字母，数字，下划线，中划线。 **取值范围**：不涉及。 **默认取值**：不涉及。
	DatasetType *string `json:"dataset_type,omitempty"`

	// **参数解释**：数据集模态细分值，如语义分割。 **约束限制**：1-40位字符。 **取值范围**：不涉及。 **默认取值**：不涉及。
	DatasetSubType *string `json:"dataset_sub_type,omitempty"`

	// **参数解释**：数据集逻辑规范，如盘古格式。 **约束限制**：1-40位字符，英文字母，数字，下划线，中划线。 **取值范围**：不涉及。 **默认取值**：不涉及。
	DatasetSchemaFormat *string `json:"dataset_schema_format,omitempty"`

	// **参数解释**：资产创建者。 **约束限制**：1-512位字符。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Creator *string `json:"creator,omitempty"`

	// **参数解释**：数据源的ID信息。 **约束限制**：1-40位字符。 **取值范围**：不涉及。 **默认取值**：不涉及。
	DataSourceId *string `json:"data_source_id,omitempty"`
}

func (o ListAssetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAssetRequest struct{}"
	}

	return strings.Join([]string{"ListAssetRequest", string(data)}, " ")
}
