package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Update2dModelTrainingJobReq 创建或更新分身数字人模型训练任务请求。
type Update2dModelTrainingJobReq struct {

	// 分身数字人模型名称。该名称会作为资产库中分身数字人模型资产名称。
	Name *string `json:"name,omitempty"`

	// 分身数字人训练任务创建者联系方式，如手机或邮箱等。
	Contact *string `json:"contact,omitempty"`

	// 命令类型： * UPDATE_VIDEO: 更新视频 * UPLOAD_VIDEO：上传视频
	CommandMessage *Update2dModelTrainingJobReqCommandMessage `json:"command_message,omitempty"`

	// 训练视频上传分片数。
	VideoMultipartCount *int32 `json:"video_multipart_count,omitempty"`

	// 分身数字人是否需要背景替换。需要背景替换的分身数字人训练视频需要绿幕拍摄。
	IsBackgroundReplacement *bool `json:"is_background_replacement,omitempty"`

	// 分身数字人训练任务的批次名称。
	BatchName *string `json:"batch_name,omitempty"`

	// 分身数字人训练任务标签。
	Tags *[]string `json:"tags,omitempty"`
}

func (o Update2dModelTrainingJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Update2dModelTrainingJobReq struct{}"
	}

	return strings.Join([]string{"Update2dModelTrainingJobReq", string(data)}, " ")
}

type Update2dModelTrainingJobReqCommandMessage struct {
	value string
}

type Update2dModelTrainingJobReqCommandMessageEnum struct {
	UPDATE_VIDEO Update2dModelTrainingJobReqCommandMessage
	UPLOAD_VIDEO Update2dModelTrainingJobReqCommandMessage
}

func GetUpdate2dModelTrainingJobReqCommandMessageEnum() Update2dModelTrainingJobReqCommandMessageEnum {
	return Update2dModelTrainingJobReqCommandMessageEnum{
		UPDATE_VIDEO: Update2dModelTrainingJobReqCommandMessage{
			value: "UPDATE_VIDEO",
		},
		UPLOAD_VIDEO: Update2dModelTrainingJobReqCommandMessage{
			value: "UPLOAD_VIDEO",
		},
	}
}

func (c Update2dModelTrainingJobReqCommandMessage) Value() string {
	return c.value
}

func (c Update2dModelTrainingJobReqCommandMessage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *Update2dModelTrainingJobReqCommandMessage) UnmarshalJSON(b []byte) error {
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
