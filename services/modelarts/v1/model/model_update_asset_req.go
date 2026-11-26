package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateAssetReq 更新资产请求体
type UpdateAssetReq struct {

	// **参数解释**：资产名称。 **约束限制**：1-128位字符，只包含中文、字母、数字、中划线、下划线。 **取值范围**：不涉及。 **默认取值**：不涉及。
	AssetName *string `json:"asset_name,omitempty"`

	// **参数解释**：资产描述。 **约束限制**：0-512位字符，不能包含英文尖括号（即 < 或 >）。 **取值范围**：不涉及。 **默认取值**：不涉及。
	AssetDesc *string `json:"asset_desc,omitempty"`

	// **参数解释**：资产组描述。 **约束限制**：0-512位字符，不能包含英文尖括号（即 < 或 >）。 **取值范围**：不涉及。 **默认取值**：不涉及。
	AssetGroupDesc *string `json:"asset_group_desc,omitempty"`

	// **参数解释**：资产品牌，例如DeepSeek,Qwen,Wan,GPT,Llama,GLM,ERNIE,Mistral,Kimi,Other。 **约束限制**：1-40位字符，仅支持英文字母、数字、下划线和中划线。 **取值范围**：不涉及。 **默认取值**：不涉及。
	AssetSeries *string `json:"asset_series,omitempty"`

	// **参数解释**：资产能力，为自由文本字符串列表(非严格枚举)。推荐使用以下约定值：例如TextGeneration,ImageUnderstanding,ImageGeneration,VideoGeneration,AgenticRL,Other。 **约束限制**：每个元素1-40位字符，仅支持英文字母、数字、下划线和中划线；最多10个元素。
	AssetCapabilities *[]string `json:"asset_capabilities,omitempty"`

	// **参数解释**：数据集资产的上下线状态。 **约束限制**：不涉及。 **取值范围**：枚举值： - online：上线 - offline：下线 - waiting: 等待中 **默认取值**：不涉及。
	Status *UpdateAssetReqStatus `json:"status,omitempty"`

	// **参数解释**：资产是否可用。 **约束限制**：不涉及。 **取值范围**： - true：可用 - false：不可用 **默认取值**：不涉及。
	Enabled *bool `json:"enabled,omitempty"`

	// **参数解释**：数据集资产是否被软删除。 **约束限制**：不涉及。 **取值范围**： - true：已删除 - false：未删除 **默认取值**：不涉及。
	IsDeleted *bool `json:"is_deleted,omitempty"`

	// **参数解释**：资产tag。 **约束限制**：不涉及。
	Tags map[string][]string `json:"tags,omitempty"`

	Properties *Properties `json:"properties,omitempty"`
}

func (o UpdateAssetReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAssetReq struct{}"
	}

	return strings.Join([]string{"UpdateAssetReq", string(data)}, " ")
}

type UpdateAssetReqStatus struct {
	value string
}

type UpdateAssetReqStatusEnum struct {
	ONLINE  UpdateAssetReqStatus
	OFFLINE UpdateAssetReqStatus
	WAITING UpdateAssetReqStatus
}

func GetUpdateAssetReqStatusEnum() UpdateAssetReqStatusEnum {
	return UpdateAssetReqStatusEnum{
		ONLINE: UpdateAssetReqStatus{
			value: "online",
		},
		OFFLINE: UpdateAssetReqStatus{
			value: "offline",
		},
		WAITING: UpdateAssetReqStatus{
			value: "waiting",
		},
	}
}

func (c UpdateAssetReqStatus) Value() string {
	return c.value
}

func (c UpdateAssetReqStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateAssetReqStatus) UnmarshalJSON(b []byte) error {
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
