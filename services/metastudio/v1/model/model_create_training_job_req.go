package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateTrainingJobReq 创建训练任务请求。
type CreateTrainingJobReq struct {
	Tag *JobTag `json:"tag,omitempty"`

	// 一段描述信息,会呈现在资产库中。
	Description *string `json:"description,omitempty"`

	// 语音性别,是男性声音还是女性声音。 * FEMALE: 女性 * MALE: 男性
	Sex *CreateTrainingJobReqSex `json:"sex,omitempty"`

	// 音色名称。该名称会作为资产库中音色模型资产名称。
	VoiceName string `json:"voice_name"`

	// 训练语言,当前仅支持中文。 * CN: 中文 * EN: 英文
	Language *string `json:"language,omitempty"`

	CreateType *CreateType `json:"create_type,omitempty"`

	// 手机号
	Phone *string `json:"phone,omitempty"`

	// 形象制作任务id
	DhtmsJobId *string `json:"dhtms_job_id,omitempty"`

	// 批次名称
	BatchName *string `json:"batch_name,omitempty"`

	// 模型输出语言类型
	OutputLanguage *string `json:"output_language,omitempty"`

	// 自定义试听文本
	CustomText *string `json:"custom_text,omitempty"`

	// 是否使用按需资源
	IsOndemandResource *bool `json:"is_ondemand_resource,omitempty"`

	// 支持的业务类型。： * VIDEO_2D：分身数字人视频制作 * LIVE_2D：分身数字人直播 * CHAT_2D：分身数字人智能交互
	SupportedService *[]SupportedServiceEnum `json:"supported_service,omitempty"`
}

func (o CreateTrainingJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTrainingJobReq struct{}"
	}

	return strings.Join([]string{"CreateTrainingJobReq", string(data)}, " ")
}

type CreateTrainingJobReqSex struct {
	value string
}

type CreateTrainingJobReqSexEnum struct {
	FEMALE CreateTrainingJobReqSex
	MALE   CreateTrainingJobReqSex
}

func GetCreateTrainingJobReqSexEnum() CreateTrainingJobReqSexEnum {
	return CreateTrainingJobReqSexEnum{
		FEMALE: CreateTrainingJobReqSex{
			value: "FEMALE",
		},
		MALE: CreateTrainingJobReqSex{
			value: "MALE",
		},
	}
}

func (c CreateTrainingJobReqSex) Value() string {
	return c.value
}

func (c CreateTrainingJobReqSex) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateTrainingJobReqSex) UnmarshalJSON(b []byte) error {
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
