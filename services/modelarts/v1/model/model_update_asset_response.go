package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateAssetResponse Response Object
type UpdateAssetResponse struct {

	// **参数解释**：资产id。 **取值范围**：不涉及。
	AssetId *string `json:"asset_id,omitempty"`

	// **参数解释**：父资产id。 **取值范围**：不涉及。
	ParentAssetId *string `json:"parent_asset_id,omitempty"`

	// **参数解释**：根资产id。 **取值范围**：不涉及。
	RootAssetId *string `json:"root_asset_id,omitempty"`

	// **参数解释**：资产API版本。 **取值范围**：v1,v2。
	ApiVersion *string `json:"api_version,omitempty"`

	// **参数解释**：资产名称。 **取值范围**：不涉及。
	AssetName *string `json:"asset_name,omitempty"`

	// **参数解释**：需要加入的资产组id。 **取值范围**：0-36位字符，已存在的组id。
	AssetGroupId *string `json:"asset_group_id,omitempty"`

	// **参数解释**：工作空间ID，[获取方法请参见[查询工作空间列表](ListWorkspace.xml)。]未创建工作空间时默认值为“0”，存在创建并使用的工作空间，以实际取值为准。 **取值范围**：1-32位字母，数字。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// **参数解释**：资产类型。 **取值范围**：枚举值： - TEMPLATE: 模板资产 - DATASET: 数据集资产 - WORKFLOW: 工作流资产 - IMAGE: 镜像资产 - ALGORITHM：算法资产 - MODEL: 模型资产 - DATASET_TEMPLATE: 数据集模板资产
	AssetType *string `json:"asset_type,omitempty"`

	// **参数解释**：资产子类型。 **取值范围**：不涉及。
	SubAssetType *string `json:"sub_asset_type,omitempty"`

	// **参数解释**：资产描述。 **取值范围**：不涉及。
	AssetDesc *string `json:"asset_desc,omitempty"`

	// **参数解释**：资产code。 **取值范围**：1-60字符，英文字母，数字，中划线，下划线，点。
	AssetCode *string `json:"asset_code,omitempty"`

	// **参数解释**：资产版本。 **取值范围**：数字或格式为x.y.z，x,y,z为数字。
	AssetVersion *string `json:"asset_version,omitempty"`

	// **参数解释**：资产tag。 **约束限制**：不涉及。
	Tags map[string][]string `json:"tags,omitempty"`

	// **参数解释**：资产支持的操作列表，例如监督微调:SFT,LORA微调:LORA,在线部署:ONLINE-DEPLOY,强化学习:AgenticRL。 **约束限制**：不涉及。
	SupportActions *interface{} `json:"support_actions,omitempty"`

	// **参数解释**：资产文件存储地址,支持OBS（对象存储服务）协议的URL和相对路径。 **取值范围**：0-1024位字符。
	AssetLocation *string `json:"asset_location,omitempty"`

	// **参数解释**：资产可执行操作详情配置。 **约束限制**：不涉及。
	Actions *[]AssetActionRsp `json:"actions,omitempty"`

	// **参数解释**：资产状态。 **取值范围**：枚举值: - online: 上线 - offline: 下线 - waiting: 等待中
	Status *UpdateAssetResponseStatus `json:"status,omitempty"`

	// **参数解释**：数据集用途。 **取值范围**：1-40位字符。
	DatasetUsage *string `json:"dataset_usage,omitempty"`

	// **参数解释**：数据集存储格式，例如JSONL，TXT，MOBI，EPUB，DOCX，PDF。 **约束限制**：1-40位字符。
	DatasetFileFormats *[]string `json:"dataset_file_formats,omitempty"`

	// **参数解释**：数据集模态划分，如音频，图像，视频。 **取值范围**：1-40位字符。
	DatasetType *string `json:"dataset_type,omitempty"`

	// **参数解释**：数据集模态细分值，如语义分割。 **取值范围**：1-40位字符。
	DatasetSubType *string `json:"dataset_sub_type,omitempty"`

	// **参数解释**：数据集逻辑规范，如盘古格式。 **取值范围**：1-40位字符。
	DatasetSchemaFormat *string `json:"dataset_schema_format,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o UpdateAssetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAssetResponse struct{}"
	}

	return strings.Join([]string{"UpdateAssetResponse", string(data)}, " ")
}

type UpdateAssetResponseStatus struct {
	value string
}

type UpdateAssetResponseStatusEnum struct {
	ONLINE  UpdateAssetResponseStatus
	OFFLINE UpdateAssetResponseStatus
	WAITING UpdateAssetResponseStatus
}

func GetUpdateAssetResponseStatusEnum() UpdateAssetResponseStatusEnum {
	return UpdateAssetResponseStatusEnum{
		ONLINE: UpdateAssetResponseStatus{
			value: "online",
		},
		OFFLINE: UpdateAssetResponseStatus{
			value: "offline",
		},
		WAITING: UpdateAssetResponseStatus{
			value: "waiting",
		},
	}
}

func (c UpdateAssetResponseStatus) Value() string {
	return c.value
}

func (c UpdateAssetResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateAssetResponseStatus) UnmarshalJSON(b []byte) error {
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
