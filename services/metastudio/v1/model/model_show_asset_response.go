package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowAssetResponse Response Object
type ShowAssetResponse struct {

	// 资产ID。
	AssetId *string `json:"asset_id,omitempty"`

	// 资产名称。
	AssetName *string `json:"asset_name,omitempty"`

	// 资产描述。
	AssetDescription *string `json:"asset_description,omitempty"`

	// 资产创建时间。
	CreateTime *string `json:"create_time,omitempty"`

	// 资产更新时间。
	UpdateTime *string `json:"update_time,omitempty"`

	// 资产类型。 * HUMAN_MODEL：数字人模型 * VOICE_MODEL：音色模型 * SCENE：场景模型 * ANIMATION：动作动画 * VIDEO：视频文件 * IMAGE：图片文件 * PPT：幻灯片文件 * MATERIAL：风格化素材 * NORMAL_MODEL: 普通模型 * COMMON_FILE：通用文件 * HUMAN_MODEL_2D:2D数字人网络模型 * BUSINESS_CARD_TEMPLET: 数字人名片模板
	AssetType *ShowAssetResponseAssetType `json:"asset_type,omitempty"`

	// 资产状态。 * CREATING：资产创建中，主文件尚未上传 * FAILED：主文件上传失败 * UNACTIVED：主文件上传成功，资产未激活，资产不可用于其他业务（用户可更新状态） * ACTIVED：主文件上传成功，资产激活，资产可用于其他业务（用户可更新状态） * DELETING：资产删除中，资产不可用，资产可恢复 * DELETED：资产文件已删除，资产不可用，资产不可恢复
	AssetState *ShowAssetResponseAssetState `json:"asset_state,omitempty"`

	// 标签列表。
	Tags *[]string `json:"tags,omitempty"`

	AssetExtraMeta *AssetExtraMeta `json:"asset_extra_meta,omitempty"`

	// 设置系统属性。
	SystemProperties *[]SystemProperty `json:"system_properties,omitempty"`

	// 资产下的文件。
	Files          *[]AssetFileInfo `json:"files,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowAssetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAssetResponse struct{}"
	}

	return strings.Join([]string{"ShowAssetResponse", string(data)}, " ")
}

type ShowAssetResponseAssetType struct {
	value string
}

type ShowAssetResponseAssetTypeEnum struct {
	HUMAN_MODEL           ShowAssetResponseAssetType
	VOICE_MODEL           ShowAssetResponseAssetType
	SCENE                 ShowAssetResponseAssetType
	ANIMATION             ShowAssetResponseAssetType
	VIDEO                 ShowAssetResponseAssetType
	IMAGE                 ShowAssetResponseAssetType
	PPT                   ShowAssetResponseAssetType
	MATERIAL              ShowAssetResponseAssetType
	NORMAL_MODEL          ShowAssetResponseAssetType
	COMMON_FILE           ShowAssetResponseAssetType
	HUMAN_MODEL_2_D       ShowAssetResponseAssetType
	BUSINESS_CARD_TEMPLET ShowAssetResponseAssetType
}

func GetShowAssetResponseAssetTypeEnum() ShowAssetResponseAssetTypeEnum {
	return ShowAssetResponseAssetTypeEnum{
		HUMAN_MODEL: ShowAssetResponseAssetType{
			value: "HUMAN_MODEL",
		},
		VOICE_MODEL: ShowAssetResponseAssetType{
			value: "VOICE_MODEL",
		},
		SCENE: ShowAssetResponseAssetType{
			value: "SCENE",
		},
		ANIMATION: ShowAssetResponseAssetType{
			value: "ANIMATION",
		},
		VIDEO: ShowAssetResponseAssetType{
			value: "VIDEO",
		},
		IMAGE: ShowAssetResponseAssetType{
			value: "IMAGE",
		},
		PPT: ShowAssetResponseAssetType{
			value: "PPT",
		},
		MATERIAL: ShowAssetResponseAssetType{
			value: "MATERIAL",
		},
		NORMAL_MODEL: ShowAssetResponseAssetType{
			value: "NORMAL_MODEL",
		},
		COMMON_FILE: ShowAssetResponseAssetType{
			value: "COMMON_FILE",
		},
		HUMAN_MODEL_2_D: ShowAssetResponseAssetType{
			value: "HUMAN_MODEL_2D",
		},
		BUSINESS_CARD_TEMPLET: ShowAssetResponseAssetType{
			value: "BUSINESS_CARD_TEMPLET",
		},
	}
}

func (c ShowAssetResponseAssetType) Value() string {
	return c.value
}

func (c ShowAssetResponseAssetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAssetResponseAssetType) UnmarshalJSON(b []byte) error {
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

type ShowAssetResponseAssetState struct {
	value string
}

type ShowAssetResponseAssetStateEnum struct {
	CREATING  ShowAssetResponseAssetState
	FAILED    ShowAssetResponseAssetState
	UNACTIVED ShowAssetResponseAssetState
	ACTIVED   ShowAssetResponseAssetState
	DELETING  ShowAssetResponseAssetState
	DELETED   ShowAssetResponseAssetState
}

func GetShowAssetResponseAssetStateEnum() ShowAssetResponseAssetStateEnum {
	return ShowAssetResponseAssetStateEnum{
		CREATING: ShowAssetResponseAssetState{
			value: "CREATING",
		},
		FAILED: ShowAssetResponseAssetState{
			value: "FAILED",
		},
		UNACTIVED: ShowAssetResponseAssetState{
			value: "UNACTIVED",
		},
		ACTIVED: ShowAssetResponseAssetState{
			value: "ACTIVED",
		},
		DELETING: ShowAssetResponseAssetState{
			value: "DELETING",
		},
		DELETED: ShowAssetResponseAssetState{
			value: "DELETED",
		},
	}
}

func (c ShowAssetResponseAssetState) Value() string {
	return c.value
}

func (c ShowAssetResponseAssetState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAssetResponseAssetState) UnmarshalJSON(b []byte) error {
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
