package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// HumanModelAssetMeta 数字人模型元数据。
type HumanModelAssetMeta struct {

	// **参数解释**： 数字人模型风格ID。 **约束限制**： 当前仅支持系统提供的几种风格。 **取值范围**： * system_male_001：男性风格01 * system_female_001：女性风格01 * system_male_002：男性风格02  * system_female_002：女性风格02  **默认取值**： 不涉及
	StyleId *string `json:"style_id,omitempty"`

	// **参数解释**： 数字人模型建模类型。 **约束限制**： 不涉及 **取值范围**： * UPLOADED：租户上传的模型 * PICTURE_MODELING：照片建模生成的模型 * CHARACTER_CUSTOMIZATION_MODELING：捏脸生成的模型
	ModelingType *HumanModelAssetMetaModelingType `json:"modeling_type,omitempty"`

	// **参数解释**： 建模任务ID。 **约束限制**： 仅由MetaStudio服务生成的模型系统自行填写 **取值范围**： 字符长度0-64位。 **默认取值**： 不涉及。
	ModelingJobId *string `json:"modeling_job_id,omitempty"`

	ModelProperties *HumanModelMetaProperties `json:"model_properties,omitempty"`

	// 可替换组件列表。
	Components *[]ComponentInfo `json:"components,omitempty"`
}

func (o HumanModelAssetMeta) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HumanModelAssetMeta struct{}"
	}

	return strings.Join([]string{"HumanModelAssetMeta", string(data)}, " ")
}

type HumanModelAssetMetaModelingType struct {
	value string
}

type HumanModelAssetMetaModelingTypeEnum struct {
	UPLOADED                         HumanModelAssetMetaModelingType
	PICTURE_MODELING                 HumanModelAssetMetaModelingType
	CHARACTER_CUSTOMIZATION_MODELING HumanModelAssetMetaModelingType
}

func GetHumanModelAssetMetaModelingTypeEnum() HumanModelAssetMetaModelingTypeEnum {
	return HumanModelAssetMetaModelingTypeEnum{
		UPLOADED: HumanModelAssetMetaModelingType{
			value: "UPLOADED",
		},
		PICTURE_MODELING: HumanModelAssetMetaModelingType{
			value: "PICTURE_MODELING",
		},
		CHARACTER_CUSTOMIZATION_MODELING: HumanModelAssetMetaModelingType{
			value: "CHARACTER_CUSTOMIZATION_MODELING",
		},
	}
}

func (c HumanModelAssetMetaModelingType) Value() string {
	return c.value
}

func (c HumanModelAssetMetaModelingType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HumanModelAssetMetaModelingType) UnmarshalJSON(b []byte) error {
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
