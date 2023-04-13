package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 更新资产请求体。
type UpdateDigitalAssetRequestBody struct {

	// 资产名称。
	AssetName *string `json:"asset_name,omitempty"`

	// 资产描述。
	AssetDescription *string `json:"asset_description,omitempty"`

	// 资产类型。 * HUMAN_MODEL:数字人模型 * VOICE_MODEL:音色模型(仅系统管理员可更新) * SCENE:场景模型 * ANIMATION:动作动画 * VIDEO:视频文件 * IMAGE:图片文件 * PPT:幻灯片文件 * MATERIAL:风格化素材
	AssetType *UpdateDigitalAssetRequestBodyAssetType `json:"asset_type,omitempty"`

	// 资产状态。 * UNACTIVED: 取消激活。未激活的资产不可用于其他业务 * ACTIVED: 激活。激活后的资产科用于其他业务
	AssetState *UpdateDigitalAssetRequestBodyAssetState `json:"asset_state,omitempty"`

	// 资产所属者。填租户的project id。
	AssetOwner *string `json:"asset_owner,omitempty"`

	// 标签列表。
	Tags *[]string `json:"tags,omitempty"`

	AssetExtraMeta *AssetExtraMeta `json:"asset_extra_meta,omitempty"`

	// 设置系统属性。
	SystemProperties *[]SystemProperty `json:"system_properties,omitempty"`
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
	HUMAN_MODEL  UpdateDigitalAssetRequestBodyAssetType
	VOICE_MODEL  UpdateDigitalAssetRequestBodyAssetType
	SCENE        UpdateDigitalAssetRequestBodyAssetType
	ANIMATION    UpdateDigitalAssetRequestBodyAssetType
	VIDEO        UpdateDigitalAssetRequestBodyAssetType
	IMAGE        UpdateDigitalAssetRequestBodyAssetType
	PPT          UpdateDigitalAssetRequestBodyAssetType
	MATERIAL     UpdateDigitalAssetRequestBodyAssetType
	NORMAL_MODEL UpdateDigitalAssetRequestBodyAssetType
	COMMON_FILE  UpdateDigitalAssetRequestBodyAssetType
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
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
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
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
