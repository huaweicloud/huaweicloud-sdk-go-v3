package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateDigitalAssetRequestBody 创建资产请求体。
type CreateDigitalAssetRequestBody struct {

	// **参数解释**： 资产名称。 **约束限制**： 不涉及。 **取值范围**： 只能使用中英文字符，字符长度0-256位。 **默认取值**： 不涉及。
	AssetName string `json:"asset_name"`

	// **参数解释**： 资产描述。 **约束限制**： 不涉及。 **取值范围**： 只能使用中英文字符，字符长度0-4096位。 **默认取值**： 不涉及。
	AssetDescription *string `json:"asset_description,omitempty"`

	// **参数解释**： 资产类型。 **约束限制**： VOICE_MODEL，HUMAN_MODEL_2D 普通用户均无法上传。 **取值范围**： 公共资产类型： * VOICE_MODEL：音色模型 * VIDEO：视频文件 * IMAGE：图片文件 * PPT：幻灯片文件 * MUSIC: 音乐 * AUDIO: 音频 * COMMON_FILE：通用文件  分身数字人资产： * HUMAN_MODEL_2D: 分身数字人模型 * BUSINESS_CARD_TEMPLET: 数字人名片模板  3D数字人资产： * HUMAN_MODEL：3D数字人模型 * SCENE：场景模型 * ANIMATION：动作动画 * MATERIAL：风格化素材 * NORMAL_MODEL: 普通模型。  **默认取值**： 不涉及。
	AssetType CreateDigitalAssetRequestBodyAssetType `json:"asset_type"`

	ReviewConfig *ReviewConfig `json:"review_config,omitempty"`

	// **参数解释**： 标签列表。 > 分身形象系统资产的tag定义如下： > - 行业：NEWS,BUSINESS,E_COMMERCE,MARKETING,KNOWLEDGE,EDUCATION,MEDICAL,SPORTS > - 性别：MALE,FEMALE > - 姿势：FULL_BODY,HALF_BODY,STANDING,SITTING,FRONT_PHOTO,SIDE_PHOTO > - 区域：ASIAN,WESTERN,MIDDLE_EASTERNER,AFRICAN,LATINO  **约束限制**： 不涉及 **取值范围**： 标签个数最大为50个。 标签内容为中英文，字符长度0-128位。 **默认取值**： 不涉及
	Tags *[]string `json:"tags,omitempty"`

	AssetExtraMeta *AssetExtraMeta `json:"asset_extra_meta,omitempty"`

	SharedConfig *AssetSharedConfig `json:"shared_config,omitempty"`

	// **参数解释**： 是否需要生成封面。 **约束限制**： 仅用于视频类资产。 **取值范围**： * true：自动生成封面。 * false：不自动生成封面。
	IsNeedGenerateCover *bool `json:"is_need_generate_cover,omitempty"`

	// **参数解释**： 用于console控制台展示顺序。 如果取值相同，则默认最新的排在前面。 **约束限制**： 不涉及 **默认取值**： 不涉及
	AssetOrder *int32 `json:"asset_order,omitempty"`

	// 支持的业务类型。： * VIDEO_2D：分身数字人视频制作 * LIVE_2D：分身数字人直播 * CHAT_2D：分身数字人智能交互
	SupportedService *[]SupportedServiceEnum `json:"supported_service,omitempty"`
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
