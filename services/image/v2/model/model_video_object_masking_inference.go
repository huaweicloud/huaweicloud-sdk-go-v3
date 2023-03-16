package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type VideoObjectMaskingInference struct {

	// 透传客户信息
	PassThrough *string `json:"pass_through,omitempty"`

	// 擦除场景，可选车内或车外场景【inside,outside】
	SceneType VideoObjectMaskingInferenceSceneType `json:"scene_type"`

	// 用户自定义产物名，无此项输入时，输出路径为{output结构体中指定输出path}/{task_id}/{task_id}.mp4；有此项输入时，输出路径为{output结构体中指定输出path}/{outcome_name}.mp4，自定义产物路径最多可定义5层文件夹目录。
	OutcomeName *string `json:"outcome_name,omitempty"`
}

func (o VideoObjectMaskingInference) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoObjectMaskingInference struct{}"
	}

	return strings.Join([]string{"VideoObjectMaskingInference", string(data)}, " ")
}

type VideoObjectMaskingInferenceSceneType struct {
	value string
}

type VideoObjectMaskingInferenceSceneTypeEnum struct {
	INSIDE  VideoObjectMaskingInferenceSceneType
	OUTSIDE VideoObjectMaskingInferenceSceneType
}

func GetVideoObjectMaskingInferenceSceneTypeEnum() VideoObjectMaskingInferenceSceneTypeEnum {
	return VideoObjectMaskingInferenceSceneTypeEnum{
		INSIDE: VideoObjectMaskingInferenceSceneType{
			value: "inside",
		},
		OUTSIDE: VideoObjectMaskingInferenceSceneType{
			value: "outside",
		},
	}
}

func (c VideoObjectMaskingInferenceSceneType) Value() string {
	return c.value
}

func (c VideoObjectMaskingInferenceSceneType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VideoObjectMaskingInferenceSceneType) UnmarshalJSON(b []byte) error {
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
