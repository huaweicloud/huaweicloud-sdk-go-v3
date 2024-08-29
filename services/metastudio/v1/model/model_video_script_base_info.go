package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type VideoScriptBaseInfo struct {

	// 剧本ID。
	ScriptId string `json:"script_id"`

	// 剧本名称。
	ScriptName string `json:"script_name"`

	// 剧本描述。
	ScriptDescription *string `json:"script_description,omitempty"`

	// 数字人模型资产ID，可以从资产库中查询。
	ModelAssetId *string `json:"model_asset_id,omitempty"`

	// 数字人模型类型。  * HUMAN_MODEL_2D：分身数字人 * HUMAN_MODEL_3D：3D数字人
	ModelAssetType *VideoScriptBaseInfoModelAssetType `json:"model_asset_type,omitempty"`

	// 剧本封面下载url。
	ScriptCoverUrl *string `json:"script_cover_url,omitempty"`

	// 脚本类型，即视频制作的驱动方式。默认TEXT * TEXT: 文本驱动，即通过TTS合成语音 * AUDIO: 语音驱动
	ScriptType *string `json:"script_type,omitempty"`

	// 台词脚本。
	Text *string `json:"text,omitempty"`

	// 创建时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o VideoScriptBaseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoScriptBaseInfo struct{}"
	}

	return strings.Join([]string{"VideoScriptBaseInfo", string(data)}, " ")
}

type VideoScriptBaseInfoModelAssetType struct {
	value string
}

type VideoScriptBaseInfoModelAssetTypeEnum struct {
	HUMAN_MODEL_2_D VideoScriptBaseInfoModelAssetType
	HUMAN_MODEL_3_D VideoScriptBaseInfoModelAssetType
}

func GetVideoScriptBaseInfoModelAssetTypeEnum() VideoScriptBaseInfoModelAssetTypeEnum {
	return VideoScriptBaseInfoModelAssetTypeEnum{
		HUMAN_MODEL_2_D: VideoScriptBaseInfoModelAssetType{
			value: "HUMAN_MODEL_2D",
		},
		HUMAN_MODEL_3_D: VideoScriptBaseInfoModelAssetType{
			value: "HUMAN_MODEL_3D",
		},
	}
}

func (c VideoScriptBaseInfoModelAssetType) Value() string {
	return c.value
}

func (c VideoScriptBaseInfoModelAssetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VideoScriptBaseInfoModelAssetType) UnmarshalJSON(b []byte) error {
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
