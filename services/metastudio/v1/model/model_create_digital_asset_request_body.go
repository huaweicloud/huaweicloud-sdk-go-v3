package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateDigitalAssetRequestBody 创建资产请求体。
type CreateDigitalAssetRequestBody struct {

	// 资产名称。
	AssetName string `json:"asset_name"`

	// 资产描述。
	AssetDescription *string `json:"asset_description,omitempty"`

	// 资产类型。  公共资产类型： * VOICE_MODEL：音色模型（仅系统管理员可上传，普通租户仅可查询） * VIDEO：视频文件 * IMAGE：图片文件 * PPT：幻灯片文件 * MUSIC: 音乐 * AUDIO: 音频 * COMMON_FILE：通用文件  分身数字人资产： * HUMAN_MODEL_2D: 分身数字人模型 * BUSINESS_CARD_TEMPLET: 数字人名片模板  3D数字人资产： * HUMAN_MODEL：3D数字人模型 * SCENE：场景模型 * ANIMATION：动作动画 * MATERIAL：风格化素材 * NORMAL_MODEL: 普通模型
	AssetType CreateDigitalAssetRequestBodyAssetType `json:"asset_type"`

	// 项目ID。 > * 仅管理员账号可设置此参数。
	AssetOwner *string `json:"asset_owner,omitempty"`

	ReviewConfig *ReviewConfig `json:"review_config,omitempty"`

	// 标签列表。
	Tags *[]string `json:"tags,omitempty"`

	AssetExtraMeta *AssetExtraMeta `json:"asset_extra_meta,omitempty"`

	// 设置系统属性。
	SystemProperties *[]SystemProperty `json:"system_properties,omitempty"`

	SharedConfig *AssetSharedConfig `json:"shared_config,omitempty"`

	// 是否需要生成封面。
	IsNeedGenerateCover *bool `json:"is_need_generate_cover,omitempty"`
}

func (o CreateDigitalAssetRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDigitalAssetRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDigitalAssetRequestBody", string(data)}, " ")
}

type CreateDigitalAssetRequestBodyAssetType struct {
	value string
}

type CreateDigitalAssetRequestBodyAssetTypeEnum struct {
	HUMAN_MODEL           CreateDigitalAssetRequestBodyAssetType
	VOICE_MODEL           CreateDigitalAssetRequestBodyAssetType
	SCENE                 CreateDigitalAssetRequestBodyAssetType
	ANIMATION             CreateDigitalAssetRequestBodyAssetType
	VIDEO                 CreateDigitalAssetRequestBodyAssetType
	IMAGE                 CreateDigitalAssetRequestBodyAssetType
	PPT                   CreateDigitalAssetRequestBodyAssetType
	MATERIAL              CreateDigitalAssetRequestBodyAssetType
	NORMAL_MODEL          CreateDigitalAssetRequestBodyAssetType
	COMMON_FILE           CreateDigitalAssetRequestBodyAssetType
	HUMAN_MODEL_2_D       CreateDigitalAssetRequestBodyAssetType
	BUSINESS_CARD_TEMPLET CreateDigitalAssetRequestBodyAssetType
	MUSIC                 CreateDigitalAssetRequestBodyAssetType
	AUDIO                 CreateDigitalAssetRequestBodyAssetType
}

func GetCreateDigitalAssetRequestBodyAssetTypeEnum() CreateDigitalAssetRequestBodyAssetTypeEnum {
	return CreateDigitalAssetRequestBodyAssetTypeEnum{
		HUMAN_MODEL: CreateDigitalAssetRequestBodyAssetType{
			value: "HUMAN_MODEL",
		},
		VOICE_MODEL: CreateDigitalAssetRequestBodyAssetType{
			value: "VOICE_MODEL",
		},
		SCENE: CreateDigitalAssetRequestBodyAssetType{
			value: "SCENE",
		},
		ANIMATION: CreateDigitalAssetRequestBodyAssetType{
			value: "ANIMATION",
		},
		VIDEO: CreateDigitalAssetRequestBodyAssetType{
			value: "VIDEO",
		},
		IMAGE: CreateDigitalAssetRequestBodyAssetType{
			value: "IMAGE",
		},
		PPT: CreateDigitalAssetRequestBodyAssetType{
			value: "PPT",
		},
		MATERIAL: CreateDigitalAssetRequestBodyAssetType{
			value: "MATERIAL",
		},
		NORMAL_MODEL: CreateDigitalAssetRequestBodyAssetType{
			value: "NORMAL_MODEL",
		},
		COMMON_FILE: CreateDigitalAssetRequestBodyAssetType{
			value: "COMMON_FILE",
		},
		HUMAN_MODEL_2_D: CreateDigitalAssetRequestBodyAssetType{
			value: "HUMAN_MODEL_2D",
		},
		BUSINESS_CARD_TEMPLET: CreateDigitalAssetRequestBodyAssetType{
			value: "BUSINESS_CARD_TEMPLET",
		},
		MUSIC: CreateDigitalAssetRequestBodyAssetType{
			value: "MUSIC",
		},
		AUDIO: CreateDigitalAssetRequestBodyAssetType{
			value: "AUDIO",
		},
	}
}

func (c CreateDigitalAssetRequestBodyAssetType) Value() string {
	return c.value
}

func (c CreateDigitalAssetRequestBodyAssetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDigitalAssetRequestBodyAssetType) UnmarshalJSON(b []byte) error {
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
