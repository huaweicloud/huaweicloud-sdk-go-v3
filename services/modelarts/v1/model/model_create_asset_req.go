package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateAssetReq **参数解释**：创建模型资产请求体。
type CreateAssetReq struct {

	// **参数解释**：父资产id。 **约束限制**：1-36位字符，英文字母，数字，中划线。 **取值范围**：不涉及。 **默认取值**：不涉及。
	ParentAssetId *string `json:"parent_asset_id,omitempty"`

	// **参数解释**：资产API版本，固定v2。 **约束限制**：固定v2。 **取值范围**：v2。 **默认取值**：v2。
	ApiVersion *string `json:"api_version,omitempty"`

	// **参数解释**：资产名称。 **约束限制**：1-128位字符，只包含中文、字母、数字、中划线、下划线。 **取值范围**：不涉及。 **默认取值**：不涉及。
	AssetName *string `json:"asset_name,omitempty"`

	// **参数解释**：用户资产来源。 **约束限制**：不涉及。 **取值范围**：枚举值： - Publish：发布 - Import：导入 **默认取值**：不涉及。
	AssetSource *CreateAssetReqAssetSource `json:"asset_source,omitempty"`

	// **参数解释**：用户资产来源类型。 **约束限制**：不涉及。 **取值范围**：枚举值： - preset_train：预训练 - custom_train：用户训练 - AgenticRL：强化学习 - DATA_CLEANING：数据清洗 - DATA_IMPORT：数据导入 - DATA_ANNOTATION：数据标注 **默认取值**：不涉及。
	AssetSourceType *CreateAssetReqAssetSourceType `json:"asset_source_type,omitempty"`

	// **参数解释**：资产code。 **约束限制**：1-60字符，英文字母，数字，中划线，下划线，点。 **取值范围**：不涉及。 **默认取值**：不涉及。
	AssetCode *string `json:"asset_code,omitempty"`

	// **参数解释**：资产描述。 **约束限制**：0-512位字符，不能包含英文尖括号（即 < 或 >）。 **取值范围**：不涉及。 **默认取值**：不涉及。
	AssetDesc *string `json:"asset_desc,omitempty"`

	// **参数解释**：资产组描述。 **约束限制**：0-512位字符，不能包含英文尖括号（即 < 或 >）。 **取值范围**：不涉及。 **默认取值**：不涉及。
	AssetGroupDesc *string `json:"asset_group_desc,omitempty"`

	// **参数解释**：资产版本号。 **约束限制**：0-99数字字符。 **取值范围**：不涉及。 **默认取值**：不涉及。
	AssetVersion string `json:"asset_version"`

	// **参数解释**：资产支持的操作列表，例如监督微调:SFT,LORA微调:LORA,在线部署:ONLINE-DEPLOY,强化学习:AgenticRL。 **约束限制**：不涉及。
	SupportActions *[]string `json:"support_actions,omitempty"`

	// **参数解释**：资产文件存储地址,支持OBS（对象存储服务）协议的URL和相对路径。 **约束限制**：0-1024位字符。 **取值范围**：不涉及。 **默认取值**：不涉及。
	AssetLocation *string `json:"asset_location,omitempty"`

	// **参数解释**：资产类型。 **约束限制**：当不传parent_asset_id时必填。 **取值范围**：枚举值: - MODEL: 模型资产 - TEMPLATE: 模板资产 - DATASET: 数据集资产 - WORKFLOW: 工作流资产 - IMAGE: 镜像资产 - ALGORITHM：算法资产 - DATASET_TEMPLATE: 数据集模板资产 - DATA_OPERATOR：算子镜像资产 **默认取值**：不涉及。
	AssetType *CreateAssetReqAssetType `json:"asset_type,omitempty"`

	// **参数解释**：需要加入的资产组id。 **约束限制**：1-36位字符，已存在的组id。 **取值范围**：不涉及。 **默认取值**：不涉及。
	AssetGroupId *string `json:"asset_group_id,omitempty"`

	// **参数解释**：资产品牌，例如DeepSeek,Qwen,Wan,GPT,Llama,GLM,ERNIE,Mistral,Kimi,Other。 **约束限制**：1-40位字符，仅支持英文字母、数字、下划线和中划线。 **取值范围**：不涉及。 **默认取值**：不涉及。
	AssetSeries *string `json:"asset_series,omitempty"`

	// **参数解释**：资产能力，为自由文本字符串列表(非严格枚举)。推荐使用以下约定值：例如TextGeneration,ImageUnderstanding,ImageGeneration,VideoGeneration,Other。 **约束限制**：每个元素1-40位字符，仅支持英文字母、数字、下划线和中划线；最多10个元素。
	AssetCapabilities *[]string `json:"asset_capabilities,omitempty"`

	// **参数解释**：资产可执行操作详情配置。 **约束限制**：不涉及。
	Actions *[]AssetActionReq `json:"actions,omitempty"`

	// **参数解释**：模型来源于哪个操作。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	SourceOperation *string `json:"source_operation,omitempty"`

	// **参数解释**：模型来源作业的名称。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	SourceJobName *string `json:"source_job_name,omitempty"`

	// **参数解释**：模型来源作业的ID。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	SourceJobId *string `json:"source_job_id,omitempty"`

	// **参数解释**：资产状态。 **约束限制**：不涉及。 **取值范围**：枚举值: - online: 上线 - offline: 下线 - waiting: 等待中 **默认取值**：不涉及。
	Status *CreateAssetReqStatus `json:"status,omitempty"`

	// **参数解释**：数据集用途。 **约束限制**：1-40位字符。 **取值范围**：不涉及。 **默认取值**：不涉及。
	DatasetUsage *string `json:"dataset_usage,omitempty"`

	// **参数解释**：数据集存储格式，例如JSONL，TXT，MOBI，EPUB，DOCX，PDF。 **约束限制**：0-40位字符。
	DatasetFileFormats *[]string `json:"dataset_file_formats,omitempty"`

	// **参数解释**：数据集模态划分，如音频，图像，视频。 **约束限制**：1-40位字符。 **取值范围**：不涉及。 **默认取值**：不涉及。
	DatasetType *string `json:"dataset_type,omitempty"`

	// **参数解释**：数据集模态细分值，如语义分割。 **约束限制**：1-40位字符。 **取值范围**：不涉及。 **默认取值**：不涉及。
	DatasetSubType *string `json:"dataset_sub_type,omitempty"`

	// **参数解释**：数据集逻辑规范，如盘古格式。 **约束限制**：1-40位字符。 **取值范围**：不涉及。 **默认取值**：不涉及。
	DatasetSchemaFormat *string `json:"dataset_schema_format,omitempty"`

	// **参数解释**：导入文件来源。 **约束限制**：数据集资产导入时需要。 **取值范围**：枚举值: - OBS: 从OBS导入 - LOCAL: 从本地导入 **默认取值**：不涉及。
	ImportFileSource *CreateAssetReqImportFileSource `json:"import_file_source,omitempty"`

	// **参数解释**：数据集资产数据源的ID。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	DataSourceId *string `json:"data_source_id,omitempty"`
}

func (o CreateAssetReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAssetReq struct{}"
	}

	return strings.Join([]string{"CreateAssetReq", string(data)}, " ")
}

type CreateAssetReqAssetSource struct {
	value string
}

type CreateAssetReqAssetSourceEnum struct {
	PUBLISH CreateAssetReqAssetSource
	IMPORT  CreateAssetReqAssetSource
}

func GetCreateAssetReqAssetSourceEnum() CreateAssetReqAssetSourceEnum {
	return CreateAssetReqAssetSourceEnum{
		PUBLISH: CreateAssetReqAssetSource{
			value: "Publish",
		},
		IMPORT: CreateAssetReqAssetSource{
			value: "Import",
		},
	}
}

func (c CreateAssetReqAssetSource) Value() string {
	return c.value
}

func (c CreateAssetReqAssetSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAssetReqAssetSource) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type CreateAssetReqAssetSourceType struct {
	value string
}

type CreateAssetReqAssetSourceTypeEnum struct {
	PRESET_TRAIN    CreateAssetReqAssetSourceType
	CUSTOM_TRAIN    CreateAssetReqAssetSourceType
	AGENTIC_RL      CreateAssetReqAssetSourceType
	DATA_CLEANING   CreateAssetReqAssetSourceType
	DATA_IMPORT     CreateAssetReqAssetSourceType
	DATA_ANNOTATION CreateAssetReqAssetSourceType
}

func GetCreateAssetReqAssetSourceTypeEnum() CreateAssetReqAssetSourceTypeEnum {
	return CreateAssetReqAssetSourceTypeEnum{
		PRESET_TRAIN: CreateAssetReqAssetSourceType{
			value: "preset_train",
		},
		CUSTOM_TRAIN: CreateAssetReqAssetSourceType{
			value: "custom_train",
		},
		AGENTIC_RL: CreateAssetReqAssetSourceType{
			value: "AgenticRL",
		},
		DATA_CLEANING: CreateAssetReqAssetSourceType{
			value: "DATA_CLEANING",
		},
		DATA_IMPORT: CreateAssetReqAssetSourceType{
			value: "DATA_IMPORT",
		},
		DATA_ANNOTATION: CreateAssetReqAssetSourceType{
			value: "DATA_ANNOTATION",
		},
	}
}

func (c CreateAssetReqAssetSourceType) Value() string {
	return c.value
}

func (c CreateAssetReqAssetSourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAssetReqAssetSourceType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type CreateAssetReqAssetType struct {
	value string
}

type CreateAssetReqAssetTypeEnum struct {
	MODEL            CreateAssetReqAssetType
	DATASET          CreateAssetReqAssetType
	TEMPLATE         CreateAssetReqAssetType
	WORKFLOW         CreateAssetReqAssetType
	IMAGE            CreateAssetReqAssetType
	ALGORITHM        CreateAssetReqAssetType
	DATASET_TEMPLATE CreateAssetReqAssetType
	DATA_OPERATOR    CreateAssetReqAssetType
}

func GetCreateAssetReqAssetTypeEnum() CreateAssetReqAssetTypeEnum {
	return CreateAssetReqAssetTypeEnum{
		MODEL: CreateAssetReqAssetType{
			value: "MODEL",
		},
		DATASET: CreateAssetReqAssetType{
			value: "DATASET",
		},
		TEMPLATE: CreateAssetReqAssetType{
			value: "TEMPLATE",
		},
		WORKFLOW: CreateAssetReqAssetType{
			value: "WORKFLOW",
		},
		IMAGE: CreateAssetReqAssetType{
			value: "IMAGE",
		},
		ALGORITHM: CreateAssetReqAssetType{
			value: "ALGORITHM",
		},
		DATASET_TEMPLATE: CreateAssetReqAssetType{
			value: "DATASET_TEMPLATE",
		},
		DATA_OPERATOR: CreateAssetReqAssetType{
			value: "DATA_OPERATOR",
		},
	}
}

func (c CreateAssetReqAssetType) Value() string {
	return c.value
}

func (c CreateAssetReqAssetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAssetReqAssetType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type CreateAssetReqStatus struct {
	value string
}

type CreateAssetReqStatusEnum struct {
	ONLINE  CreateAssetReqStatus
	OFFLINE CreateAssetReqStatus
	WAITING CreateAssetReqStatus
}

func GetCreateAssetReqStatusEnum() CreateAssetReqStatusEnum {
	return CreateAssetReqStatusEnum{
		ONLINE: CreateAssetReqStatus{
			value: "online",
		},
		OFFLINE: CreateAssetReqStatus{
			value: "offline",
		},
		WAITING: CreateAssetReqStatus{
			value: "waiting",
		},
	}
}

func (c CreateAssetReqStatus) Value() string {
	return c.value
}

func (c CreateAssetReqStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAssetReqStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type CreateAssetReqImportFileSource struct {
	value string
}

type CreateAssetReqImportFileSourceEnum struct {
	OBS        CreateAssetReqImportFileSource
	LOCAL      CreateAssetReqImportFileSource
	AIDATALAKE CreateAssetReqImportFileSource
}

func GetCreateAssetReqImportFileSourceEnum() CreateAssetReqImportFileSourceEnum {
	return CreateAssetReqImportFileSourceEnum{
		OBS: CreateAssetReqImportFileSource{
			value: "OBS",
		},
		LOCAL: CreateAssetReqImportFileSource{
			value: "LOCAL",
		},
		AIDATALAKE: CreateAssetReqImportFileSource{
			value: "AIDATALAKE",
		},
	}
}

func (c CreateAssetReqImportFileSource) Value() string {
	return c.value
}

func (c CreateAssetReqImportFileSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAssetReqImportFileSource) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
