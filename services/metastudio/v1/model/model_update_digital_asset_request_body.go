package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateDigitalAssetRequestBody 更新资产请求体。
type UpdateDigitalAssetRequestBody struct {

	// **参数解释**： 资产名称。 **约束限制**： 不涉及。 **取值范围**： 只能使用中英文字符，字符长度0-256位。 **默认取值**： 不涉及。
	AssetName *string `json:"asset_name,omitempty"`

	// **参数解释**： 资产描述。 **约束限制**： 不涉及。 **取值范围**： 只能使用中英文字符，字符长度0-4096位。 **默认取值**： 不涉及。
	AssetDescription *string `json:"asset_description,omitempty"`

	// **参数解释**： 资产类型。 **约束限制**： VOICE_MODEL，HUMAN_MODEL_2D 普通用户均无法上传。 **取值范围**： 公共资产类型： * VOICE_MODEL：音色模型 * VIDEO：视频文件 * IMAGE：图片文件 * PPT：幻灯片文件 * MUSIC: 音乐 * AUDIO: 音频 * COMMON_FILE：通用文件  分身数字人资产： * HUMAN_MODEL_2D: 分身数字人模型 * BUSINESS_CARD_TEMPLET: 数字人名片模板  3D数字人资产： * HUMAN_MODEL：3D数字人模型 * SCENE：场景模型 * ANIMATION：动作动画 * MATERIAL：风格化素材 * NORMAL_MODEL: 普通模型。  **默认取值**： 不涉及。
	AssetType *UpdateDigitalAssetRequestBodyAssetType `json:"asset_type,omitempty"`

	// **参数解释**： 资产状态。 **约束限制**： 租户仅能激活或取消激活资产，其他状态由系统自动更新。 **取值范围**： * UNACTIVED：取消激活。未激活的资产不可用于其他业务 * ACTIVED：激活。激活后的资产可用于其他业务 * WAITING_DELETE：资产将被下线(激活状态资产可用、管理员可用)  **默认取值**： 不涉及。
	AssetState *UpdateDigitalAssetRequestBodyAssetState `json:"asset_state,omitempty"`

	ReviewConfig *ReviewConfig `json:"review_config,omitempty"`

	// **参数解释**： 标签列表。 > 分身形象系统资产的tag定义如下： > - 行业：NEWS,BUSINESS,E-COMMERCE,MARKETING,KNOWLEDGE,EDUCATION,SPORTS > - 性别：MALE,FEMALE > - 姿势：FULL-BODY,HALF-BODY,STANDING,SITTING,WALKING > - 区域：ASIAN,WESTERN,MIDDLE-EASTERNER,AFRICAN,LATINO  **约束限制**： 不涉及 **取值范围**： 标签个数最大为50个。 标签内容为中英文，字符长度0-128位。 **默认取值**： 不涉及
	Tags *[]string `json:"tags,omitempty"`

	AssetExtraMeta *AssetExtraMeta `json:"asset_extra_meta,omitempty"`

	// 设置系统属性。
	SystemProperties *[]SystemProperty `json:"system_properties,omitempty"`

	SharedConfig *AssetSharedConfig `json:"shared_config,omitempty"`

	// **参数解释**： 用于console控制台展示顺序。 如果取值相同，则默认最新的排在前面。 **约束限制**： 不涉及 **默认取值**： 不涉及
	AssetOrder *int32 `json:"asset_order,omitempty"`

	// 支持的业务类型。： * VIDEO_2D：分身数字人视频制作 * LIVE_2D：分身数字人直播 * CHAT_2D：分身数字人智能交互
	SupportedService *[]SupportedServiceEnum `json:"supported_service,omitempty"`
}

func (o UpdateDigitalAssetRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDigitalAssetRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateDigitalAssetRequestBody", string(data)}, " ")
}

type UpdateDigitalAssetRequestBodyAssetType struct {
	value string
}

type UpdateDigitalAssetRequestBodyAssetTypeEnum struct {
	HUMAN_MODEL           UpdateDigitalAssetRequestBodyAssetType
	VOICE_MODEL           UpdateDigitalAssetRequestBodyAssetType
	SCENE                 UpdateDigitalAssetRequestBodyAssetType
	ANIMATION             UpdateDigitalAssetRequestBodyAssetType
	VIDEO                 UpdateDigitalAssetRequestBodyAssetType
	IMAGE                 UpdateDigitalAssetRequestBodyAssetType
	PPT                   UpdateDigitalAssetRequestBodyAssetType
	MATERIAL              UpdateDigitalAssetRequestBodyAssetType
	NORMAL_MODEL          UpdateDigitalAssetRequestBodyAssetType
	COMMON_FILE           UpdateDigitalAssetRequestBodyAssetType
	HUMAN_MODEL_2_D       UpdateDigitalAssetRequestBodyAssetType
	BUSINESS_CARD_TEMPLET UpdateDigitalAssetRequestBodyAssetType
	MUSIC                 UpdateDigitalAssetRequestBodyAssetType
	AUDIO                 UpdateDigitalAssetRequestBodyAssetType
}

func GetUpdateDigitalAssetRequestBodyAssetTypeEnum() UpdateDigitalAssetRequestBodyAssetTypeEnum {
	return UpdateDigitalAssetRequestBodyAssetTypeEnum{
		HUMAN_MODEL: UpdateDigitalAssetRequestBodyAssetType{
			value: "HUMAN_MODEL",
		},
		VOICE_MODEL: UpdateDigitalAssetRequestBodyAssetType{
			value: "VOICE_MODEL",
		},
		SCENE: UpdateDigitalAssetRequestBodyAssetType{
			value: "SCENE",
		},
		ANIMATION: UpdateDigitalAssetRequestBodyAssetType{
			value: "ANIMATION",
		},
		VIDEO: UpdateDigitalAssetRequestBodyAssetType{
			value: "VIDEO",
		},
		IMAGE: UpdateDigitalAssetRequestBodyAssetType{
			value: "IMAGE",
		},
		PPT: UpdateDigitalAssetRequestBodyAssetType{
			value: "PPT",
		},
		MATERIAL: UpdateDigitalAssetRequestBodyAssetType{
			value: "MATERIAL",
		},
		NORMAL_MODEL: UpdateDigitalAssetRequestBodyAssetType{
			value: "NORMAL_MODEL",
		},
		COMMON_FILE: UpdateDigitalAssetRequestBodyAssetType{
			value: "COMMON_FILE",
		},
		HUMAN_MODEL_2_D: UpdateDigitalAssetRequestBodyAssetType{
			value: "HUMAN_MODEL_2D",
		},
		BUSINESS_CARD_TEMPLET: UpdateDigitalAssetRequestBodyAssetType{
			value: "BUSINESS_CARD_TEMPLET",
		},
		MUSIC: UpdateDigitalAssetRequestBodyAssetType{
			value: "MUSIC",
		},
		AUDIO: UpdateDigitalAssetRequestBodyAssetType{
			value: "AUDIO",
		},
	}
}

func (c UpdateDigitalAssetRequestBodyAssetType) Value() string {
	return c.value
}

func (c UpdateDigitalAssetRequestBodyAssetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateDigitalAssetRequestBodyAssetType) UnmarshalJSON(b []byte) error {
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

type UpdateDigitalAssetRequestBodyAssetState struct {
	value string
}

type UpdateDigitalAssetRequestBodyAssetStateEnum struct {
	UNACTIVED      UpdateDigitalAssetRequestBodyAssetState
	ACTIVED        UpdateDigitalAssetRequestBodyAssetState
	WAITING_DELETE UpdateDigitalAssetRequestBodyAssetState
}

func GetUpdateDigitalAssetRequestBodyAssetStateEnum() UpdateDigitalAssetRequestBodyAssetStateEnum {
	return UpdateDigitalAssetRequestBodyAssetStateEnum{
		UNACTIVED: UpdateDigitalAssetRequestBodyAssetState{
			value: "UNACTIVED",
		},
		ACTIVED: UpdateDigitalAssetRequestBodyAssetState{
			value: "ACTIVED",
		},
		WAITING_DELETE: UpdateDigitalAssetRequestBodyAssetState{
			value: "WAITING_DELETE",
		},
	}
}

func (c UpdateDigitalAssetRequestBodyAssetState) Value() string {
	return c.value
}

func (c UpdateDigitalAssetRequestBodyAssetState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateDigitalAssetRequestBodyAssetState) UnmarshalJSON(b []byte) error {
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
