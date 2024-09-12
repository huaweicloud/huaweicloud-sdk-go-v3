package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Create2dModelTrainingJobReq 创建或更新分身数字人模型训练任务请求。
type Create2dModelTrainingJobReq struct {

	// 分身数字人模型名称。该名称会作为资产库中分身数字人模型资产名称。
	Name string `json:"name"`

	// 分身数字人训练任务创建者联系方式，如手机或邮箱等。
	Contact *string `json:"contact,omitempty"`

	// 命令类型： * UPDATE_VIDEO: 更新视频 * UPLOAD_VIDEO：上传视频 * CONFIRM_ACTION_VIDEO: 确认动作编排视频 * GET_ACTION_VIDEO_MULTIPART: 获取动作编排视频分片
	CommandMessage *Create2dModelTrainingJobReqCommandMessage `json:"command_message,omitempty"`

	// 训练视频上传分片数（上传时对唯一训练视频文件的数据分片，用于对该文件的并发上传，不是分多个视频文件上传）。
	VideoMultipartCount *int32 `json:"video_multipart_count,omitempty"`

	// 动作视频上传分片数。
	ActionVideoMultipartCount *int32 `json:"action_video_multipart_count,omitempty"`

	// 分身数字人是否需要背景替换。需要背景替换的分身数字人训练视频需要绿幕拍摄。
	IsBackgroundReplacement *bool `json:"is_background_replacement,omitempty"`

	// 分身数字人训练任务的批次名称。
	BatchName *string `json:"batch_name,omitempty"`

	// 分身数字人训练任务标签。
	Tags *[]string `json:"tags,omitempty"`

	// 分身数字人模型版本。默认是V3版本模型。 * V2: V2版本模型 * V3：V3版本模型 * V3.2：V3.2版本模型 > * V2版本已废弃不用
	ModelVersion *Create2dModelTrainingJobReqModelVersion `json:"model_version,omitempty"`

	// 是否是基础版的形象训练
	IsFlexus *bool `json:"is_flexus,omitempty"`

	// 声音来源类型 * VIDEO：视频中抽取音频 * AUDIO：单独上传的音频
	AudioSourceType *Create2dModelTrainingJobReqAudioSourceType `json:"audio_source_type,omitempty"`

	VoiceProperties *VoiceProperties `json:"voice_properties,omitempty"`

	// 该任务所生成的模型支持的业务类型，可多选
	SupportedService *[]SupportedServiceEnum `json:"supported_service,omitempty"`
}

func (o Create2dModelTrainingJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Create2dModelTrainingJobReq struct{}"
	}

	return strings.Join([]string{"Create2dModelTrainingJobReq", string(data)}, " ")
}

type Create2dModelTrainingJobReqCommandMessage struct {
	value string
}

type Create2dModelTrainingJobReqCommandMessageEnum struct {
	UPDATE_VIDEO               Create2dModelTrainingJobReqCommandMessage
	UPLOAD_VIDEO               Create2dModelTrainingJobReqCommandMessage
	CONFIRM_ACTION_VIDEO       Create2dModelTrainingJobReqCommandMessage
	GET_ACTION_VIDEO_MULTIPART Create2dModelTrainingJobReqCommandMessage
}

func GetCreate2dModelTrainingJobReqCommandMessageEnum() Create2dModelTrainingJobReqCommandMessageEnum {
	return Create2dModelTrainingJobReqCommandMessageEnum{
		UPDATE_VIDEO: Create2dModelTrainingJobReqCommandMessage{
			value: "UPDATE_VIDEO",
		},
		UPLOAD_VIDEO: Create2dModelTrainingJobReqCommandMessage{
			value: "UPLOAD_VIDEO",
		},
		CONFIRM_ACTION_VIDEO: Create2dModelTrainingJobReqCommandMessage{
			value: "CONFIRM_ACTION_VIDEO",
		},
		GET_ACTION_VIDEO_MULTIPART: Create2dModelTrainingJobReqCommandMessage{
			value: "GET_ACTION_VIDEO_MULTIPART",
		},
	}
}

func (c Create2dModelTrainingJobReqCommandMessage) Value() string {
	return c.value
}

func (c Create2dModelTrainingJobReqCommandMessage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *Create2dModelTrainingJobReqCommandMessage) UnmarshalJSON(b []byte) error {
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

type Create2dModelTrainingJobReqModelVersion struct {
	value string
}

type Create2dModelTrainingJobReqModelVersionEnum struct {
	V2   Create2dModelTrainingJobReqModelVersion
	V3   Create2dModelTrainingJobReqModelVersion
	V3_2 Create2dModelTrainingJobReqModelVersion
}

func GetCreate2dModelTrainingJobReqModelVersionEnum() Create2dModelTrainingJobReqModelVersionEnum {
	return Create2dModelTrainingJobReqModelVersionEnum{
		V2: Create2dModelTrainingJobReqModelVersion{
			value: "V2",
		},
		V3: Create2dModelTrainingJobReqModelVersion{
			value: "V3",
		},
		V3_2: Create2dModelTrainingJobReqModelVersion{
			value: "V3.2",
		},
	}
}

func (c Create2dModelTrainingJobReqModelVersion) Value() string {
	return c.value
}

func (c Create2dModelTrainingJobReqModelVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *Create2dModelTrainingJobReqModelVersion) UnmarshalJSON(b []byte) error {
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

type Create2dModelTrainingJobReqAudioSourceType struct {
	value string
}

type Create2dModelTrainingJobReqAudioSourceTypeEnum struct {
	VIDEO Create2dModelTrainingJobReqAudioSourceType
	AUDIO Create2dModelTrainingJobReqAudioSourceType
}

func GetCreate2dModelTrainingJobReqAudioSourceTypeEnum() Create2dModelTrainingJobReqAudioSourceTypeEnum {
	return Create2dModelTrainingJobReqAudioSourceTypeEnum{
		VIDEO: Create2dModelTrainingJobReqAudioSourceType{
			value: "VIDEO",
		},
		AUDIO: Create2dModelTrainingJobReqAudioSourceType{
			value: "AUDIO",
		},
	}
}

func (c Create2dModelTrainingJobReqAudioSourceType) Value() string {
	return c.value
}

func (c Create2dModelTrainingJobReqAudioSourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *Create2dModelTrainingJobReqAudioSourceType) UnmarshalJSON(b []byte) error {
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
