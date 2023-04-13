package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 照片建模任务详情。
type PictureModelingInfo struct {

	// 照片建模任务ID。
	JobId string `json:"job_id"`

	// 任务的状态。 * WAITING：等待任务调度 * PROCESSING：正在处理 * PARTIAL_SUCCEED: 部分成功（模型生成，截图失败） * SUCCEED：成功 * FAILED：失败 * CANCELED：取消
	State PictureModelingInfoState `json:"state"`

	// 任务开始时间,格式遵循：RFC 3339。 例 “2020-07-30T10:43:17Z”。
	StartTime *string `json:"start_time,omitempty"`

	// 任务结束时间,格式遵循：RFC 3339。 例 “2020-07-30T10:43:17Z”。
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
}

func (o PictureModelingInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PictureModelingInfo struct{}"
	}

	return strings.Join([]string{"PictureModelingInfo", string(data)}, " ")
}

type PictureModelingInfoState struct {
	value string
}

type PictureModelingInfoStateEnum struct {
	WAITING         PictureModelingInfoState
	PROCESSING      PictureModelingInfoState
	PARTIAL_SUCCEED PictureModelingInfoState
	SUCCEED         PictureModelingInfoState
	FAILED          PictureModelingInfoState
	CANCELED        PictureModelingInfoState
}

func GetPictureModelingInfoStateEnum() PictureModelingInfoStateEnum {
	return PictureModelingInfoStateEnum{
		WAITING: PictureModelingInfoState{
			value: "WAITING",
		},
		PROCESSING: PictureModelingInfoState{
			value: "PROCESSING",
		},
		PARTIAL_SUCCEED: PictureModelingInfoState{
			value: "PARTIAL_SUCCEED",
		},
		SUCCEED: PictureModelingInfoState{
			value: "SUCCEED",
		},
		FAILED: PictureModelingInfoState{
			value: "FAILED",
		},
		CANCELED: PictureModelingInfoState{
			value: "CANCELED",
		},
	}
}

func (c PictureModelingInfoState) Value() string {
	return c.value
}

func (c PictureModelingInfoState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PictureModelingInfoState) UnmarshalJSON(b []byte) error {
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
