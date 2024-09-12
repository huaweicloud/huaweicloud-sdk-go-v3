package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Execute2dModelTrainingCommandByUserResponse Response Object
type Execute2dModelTrainingCommandByUserResponse struct {

	// 命令执行结果。 * EXCUTE_SUCCESS: 命令提交成功 * EXCUTE_FAILED: 命令提交失败
	CommondResult *Execute2dModelTrainingCommandByUserResponseCommondResult `json:"commond_result,omitempty"`

	// 附件上传地址
	AttachmentUploadUrl *[]string `json:"attachment_upload_url,omitempty"`

	// 训练视频已上传分片信息
	MultipartData *[]MultipartUploadInfo `json:"multipart_data,omitempty"`

	// 命令执行失败原因描述
	ExcuteFailedMsg *string `json:"excute_failed_msg,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o Execute2dModelTrainingCommandByUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Execute2dModelTrainingCommandByUserResponse struct{}"
	}

	return strings.Join([]string{"Execute2dModelTrainingCommandByUserResponse", string(data)}, " ")
}

type Execute2dModelTrainingCommandByUserResponseCommondResult struct {
	value string
}

type Execute2dModelTrainingCommandByUserResponseCommondResultEnum struct {
	EXCUTE_SUCCESS Execute2dModelTrainingCommandByUserResponseCommondResult
	EXCUTE_FAILED  Execute2dModelTrainingCommandByUserResponseCommondResult
}

func GetExecute2dModelTrainingCommandByUserResponseCommondResultEnum() Execute2dModelTrainingCommandByUserResponseCommondResultEnum {
	return Execute2dModelTrainingCommandByUserResponseCommondResultEnum{
		EXCUTE_SUCCESS: Execute2dModelTrainingCommandByUserResponseCommondResult{
			value: "EXCUTE_SUCCESS",
		},
		EXCUTE_FAILED: Execute2dModelTrainingCommandByUserResponseCommondResult{
			value: "EXCUTE_FAILED",
		},
	}
}

func (c Execute2dModelTrainingCommandByUserResponseCommondResult) Value() string {
	return c.value
}

func (c Execute2dModelTrainingCommandByUserResponseCommondResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *Execute2dModelTrainingCommandByUserResponseCommondResult) UnmarshalJSON(b []byte) error {
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
