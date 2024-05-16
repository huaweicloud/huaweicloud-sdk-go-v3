package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateDigitalAssetRequestBody 更新资产请求体。
type UpdateDigitalAssetRequestBody struct {

	// 资产名称。
	AssetName *string `json:"asset_name,omitempty"`

	// 资产描述。
	AssetDescription *string `json:"asset_description,omitempty"`

	// 资产类型。  公共资产类型： * VOICE_MODEL：音色模型（仅系统管理员可上传，普通租户仅可查询） * VIDEO：视频文件 * IMAGE：图片文件 * PPT：幻灯片文件 * MUSIC: 音乐 * AUDIO: 音频 * COMMON_FILE：通用文件  分身数字人资产： * HUMAN_MODEL_2D: 分身数字人模型 * BUSINESS_CARD_TEMPLET: 数字人名片模板  3D数字人资产： * HUMAN_MODEL：3D数字人模型 * SCENE：场景模型 * ANIMATION：动作动画 * MATERIAL：风格化素材 * NORMAL_MODEL: 普通模型
	AssetType *UpdateDigitalAssetRequestBodyAssetType `json:"asset_type,omitempty"`

	// 资产状态。 * UNACTIVED：取消激活。未激活的资产不可用于其他业务 * ACTIVED：激活。激活后的资产可用于其他业务
	AssetState *UpdateDigitalAssetRequestBodyAssetState `json:"asset_state,omitempty"`

	// 项目ID。 > * 仅管理员账号可设置此参数。
	AssetOwner *string `json:"asset_owner,omitempty"`

	ReviewConfig *ReviewConfig `json:"review_config,omitempty"`

	// 标签列表。
	Tags *[]string `json:"tags,omitempty"`

	AssetExtraMeta *AssetExtraMeta `json:"asset_extra_meta,omitempty"`

	// 设置系统属性。
	SystemProperties *[]SystemProperty `json:"system_properties,omitempty"`

	SharedConfig *SharedConfig `json:"shared_config,omitempty"`
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
	UNACTIVED UpdateDigitalAssetRequestBodyAssetState
	ACTIVED   UpdateDigitalAssetRequestBodyAssetState
}

func GetUpdateDigitalAssetRequestBodyAssetStateEnum() UpdateDigitalAssetRequestBodyAssetStateEnum {
	return UpdateDigitalAssetRequestBodyAssetStateEnum{
		UNACTIVED: UpdateDigitalAssetRequestBodyAssetState{
			value: "UNACTIVED",
		},
		ACTIVED: UpdateDigitalAssetRequestBodyAssetState{
			value: "ACTIVED",
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
