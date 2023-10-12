package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowPhotoDetectionResponse Response Object
type ShowPhotoDetectionResponse struct {

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 任务的状态。 * WAITING：等待 * PROCESSING：处理中 * SUCCEED：成功 * FAILED：失败 * CANCELED：取消
	State *ShowPhotoDetectionResponseState `json:"state,omitempty"`

	ErrorInfo *ErrorResponse `json:"error_info,omitempty"`

	// 任务创建时间。
	CreateTime *string `json:"create_time,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPhotoDetectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPhotoDetectionResponse struct{}"
	}

	return strings.Join([]string{"ShowPhotoDetectionResponse", string(data)}, " ")
}

type ShowPhotoDetectionResponseState struct {
	value string
}

type ShowPhotoDetectionResponseStateEnum struct {
	WAITING    ShowPhotoDetectionResponseState
	PROCESSING ShowPhotoDetectionResponseState
	SUCCEED    ShowPhotoDetectionResponseState
	FAILED     ShowPhotoDetectionResponseState
	CANCELED   ShowPhotoDetectionResponseState
}

func GetShowPhotoDetectionResponseStateEnum() ShowPhotoDetectionResponseStateEnum {
	return ShowPhotoDetectionResponseStateEnum{
		WAITING: ShowPhotoDetectionResponseState{
			value: "WAITING",
		},
		PROCESSING: ShowPhotoDetectionResponseState{
			value: "PROCESSING",
		},
		SUCCEED: ShowPhotoDetectionResponseState{
			value: "SUCCEED",
		},
		FAILED: ShowPhotoDetectionResponseState{
			value: "FAILED",
		},
		CANCELED: ShowPhotoDetectionResponseState{
			value: "CANCELED",
		},
	}
}

func (c ShowPhotoDetectionResponseState) Value() string {
	return c.value
}

func (c ShowPhotoDetectionResponseState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowPhotoDetectionResponseState) UnmarshalJSON(b []byte) error {
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
