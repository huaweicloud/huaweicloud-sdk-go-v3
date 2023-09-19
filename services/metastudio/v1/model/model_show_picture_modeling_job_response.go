package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowPictureModelingJobResponse Response Object
type ShowPictureModelingJobResponse struct {

	// 照片建模任务ID。
	JobId string `json:"job_id"`

	// 任务的状态。 * WAITING：等待任务调度 * PROCESSING：正在处理 * PARTIAL_SUCCEED：部分成功（模型生成，截图失败） * SUCCEED：成功 * FAILED：失败 * CANCELED：取消
	State ShowPictureModelingJobResponseState `json:"state"`

	// 任务开始时间，格式遵循：RFC 3339。 例 “2020-07-30T10:43:17Z”。
	StartTime *string `json:"start_time,omitempty"`

	// 任务结束时间，格式遵循：RFC 3339。 例 “2020-07-30T10:43:17Z”。
	EndTime *string `json:"end_time,omitempty"`

	ErrorInfo *ErrorResponse `json:"error_info,omitempty"`

	// 模型资产ID。
	ModelAssetId *string `json:"model_asset_id,omitempty"`

	// 数字人模型名称。
	Name *string `json:"name,omitempty"`

	// 风格ID。
	StyleId *string `json:"style_id,omitempty"`

	// 模型封面URL。
	ModelCoverUrl *string `json:"model_cover_url,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPictureModelingJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPictureModelingJobResponse struct{}"
	}

	return strings.Join([]string{"ShowPictureModelingJobResponse", string(data)}, " ")
}

type ShowPictureModelingJobResponseState struct {
	value string
}

type ShowPictureModelingJobResponseStateEnum struct {
	WAITING         ShowPictureModelingJobResponseState
	PROCESSING      ShowPictureModelingJobResponseState
	PARTIAL_SUCCEED ShowPictureModelingJobResponseState
	SUCCEED         ShowPictureModelingJobResponseState
	FAILED          ShowPictureModelingJobResponseState
	CANCELED        ShowPictureModelingJobResponseState
}

func GetShowPictureModelingJobResponseStateEnum() ShowPictureModelingJobResponseStateEnum {
	return ShowPictureModelingJobResponseStateEnum{
		WAITING: ShowPictureModelingJobResponseState{
			value: "WAITING",
		},
		PROCESSING: ShowPictureModelingJobResponseState{
			value: "PROCESSING",
		},
		PARTIAL_SUCCEED: ShowPictureModelingJobResponseState{
			value: "PARTIAL_SUCCEED",
		},
		SUCCEED: ShowPictureModelingJobResponseState{
			value: "SUCCEED",
		},
		FAILED: ShowPictureModelingJobResponseState{
			value: "FAILED",
		},
		CANCELED: ShowPictureModelingJobResponseState{
			value: "CANCELED",
		},
	}
}

func (c ShowPictureModelingJobResponseState) Value() string {
	return c.value
}

func (c ShowPictureModelingJobResponseState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowPictureModelingJobResponseState) UnmarshalJSON(b []byte) error {
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
