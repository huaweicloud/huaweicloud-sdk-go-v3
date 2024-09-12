package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DigitalHumanVideo 数字人制作任务信息。
type DigitalHumanVideo struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 任务的状态。 * WAITING：等待 * PROCESSING：处理中 * SUCCEED：成功 * FAILED：失败 * CANCELED：取消 * BLOCK: 冻结
	State DigitalHumanVideoState `json:"state"`

	// 任务类型。 * 2D_DIGITAL_HUMAN_VIDEO: 分身数字人视频制作任务 * PHOTO_DIGITAL_HUMAN_VIDEO: 照片数字人视频制作任务
	JobType *DigitalHumanVideoJobType `json:"job_type,omitempty"`

	// 数字人视频制作开始时间。
	StartTime *string `json:"start_time,omitempty"`

	// 数字人视频制作结束时间。
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释**： 数字人视频内容时长。
	Duration *float32 `json:"duration,omitempty"`

	OutputAssetConfig *OutputAssetInfo `json:"output_asset_config,omitempty"`

	ErrorInfo *ErrorResponse `json:"error_info,omitempty"`

	// 任务创建时间。
	CreateTime *string `json:"create_time,omitempty"`

	// 任务更新时间。
	LastupdateTime *string `json:"lastupdate_time,omitempty"`
}

func (o DigitalHumanVideo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DigitalHumanVideo struct{}"
	}

	return strings.Join([]string{"DigitalHumanVideo", string(data)}, " ")
}

type DigitalHumanVideoState struct {
	value string
}

type DigitalHumanVideoStateEnum struct {
	WAITING    DigitalHumanVideoState
	PROCESSING DigitalHumanVideoState
	SUCCEED    DigitalHumanVideoState
	FAILED     DigitalHumanVideoState
	CANCELED   DigitalHumanVideoState
	BLOCK      DigitalHumanVideoState
}

func GetDigitalHumanVideoStateEnum() DigitalHumanVideoStateEnum {
	return DigitalHumanVideoStateEnum{
		WAITING: DigitalHumanVideoState{
			value: "WAITING",
		},
		PROCESSING: DigitalHumanVideoState{
			value: "PROCESSING",
		},
		SUCCEED: DigitalHumanVideoState{
			value: "SUCCEED",
		},
		FAILED: DigitalHumanVideoState{
			value: "FAILED",
		},
		CANCELED: DigitalHumanVideoState{
			value: "CANCELED",
		},
		BLOCK: DigitalHumanVideoState{
			value: "BLOCK",
		},
	}
}

func (c DigitalHumanVideoState) Value() string {
	return c.value
}

func (c DigitalHumanVideoState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DigitalHumanVideoState) UnmarshalJSON(b []byte) error {
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

type DigitalHumanVideoJobType struct {
	value string
}

type DigitalHumanVideoJobTypeEnum struct {
	E_2_D_DIGITAL_HUMAN_VIDEO DigitalHumanVideoJobType
	DIGITAL_HUMAN_PHOTO_VIDEO DigitalHumanVideoJobType
}

func GetDigitalHumanVideoJobTypeEnum() DigitalHumanVideoJobTypeEnum {
	return DigitalHumanVideoJobTypeEnum{
		E_2_D_DIGITAL_HUMAN_VIDEO: DigitalHumanVideoJobType{
			value: "2D_DIGITAL_HUMAN_VIDEO",
		},
		DIGITAL_HUMAN_PHOTO_VIDEO: DigitalHumanVideoJobType{
			value: "DIGITAL_HUMAN_PHOTO_VIDEO",
		},
	}
}

func (c DigitalHumanVideoJobType) Value() string {
	return c.value
}

func (c DigitalHumanVideoJobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DigitalHumanVideoJobType) UnmarshalJSON(b []byte) error {
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
