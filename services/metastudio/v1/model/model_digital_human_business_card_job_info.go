package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DigitalHumanBusinessCardJobInfo 数字人名片制作任务信息。
type DigitalHumanBusinessCardJobInfo struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 任务的状态。 * WAITING: 等待 * PROCESSING: 处理中 * SUCCEED: 成功 * FAILED: 失败 * CANCELED: 取消 * BLOCK: 冻结
	State DigitalHumanBusinessCardJobInfoState `json:"state"`

	// 数字人名片制作开始时间。
	StartTime *string `json:"start_time,omitempty"`

	// 数字人名片制作结束时间。
	EndTime *string `json:"end_time,omitempty"`

	OutputAssetConfig *OutputAssetInfo `json:"output_asset_config,omitempty"`

	ErrorInfo *ErrorResponse `json:"error_info,omitempty"`

	// 任务创建时间。
	CreateTime *string `json:"create_time,omitempty"`

	// 任务更新时间。
	LastupdateTime *string `json:"lastupdate_time,omitempty"`

	// 数字人名片类型。 * 2D_DIGITAL_HUMAN_CARD：分身数字人名片
	BusinessCardType *DigitalHumanBusinessCardJobInfoBusinessCardType `json:"business_card_type,omitempty"`
}

func (o DigitalHumanBusinessCardJobInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DigitalHumanBusinessCardJobInfo struct{}"
	}

	return strings.Join([]string{"DigitalHumanBusinessCardJobInfo", string(data)}, " ")
}

type DigitalHumanBusinessCardJobInfoState struct {
	value string
}

type DigitalHumanBusinessCardJobInfoStateEnum struct {
	WAITING    DigitalHumanBusinessCardJobInfoState
	PROCESSING DigitalHumanBusinessCardJobInfoState
	SUCCEED    DigitalHumanBusinessCardJobInfoState
	FAILED     DigitalHumanBusinessCardJobInfoState
	CANCELED   DigitalHumanBusinessCardJobInfoState
	BLOCK      DigitalHumanBusinessCardJobInfoState
}

func GetDigitalHumanBusinessCardJobInfoStateEnum() DigitalHumanBusinessCardJobInfoStateEnum {
	return DigitalHumanBusinessCardJobInfoStateEnum{
		WAITING: DigitalHumanBusinessCardJobInfoState{
			value: "WAITING",
		},
		PROCESSING: DigitalHumanBusinessCardJobInfoState{
			value: "PROCESSING",
		},
		SUCCEED: DigitalHumanBusinessCardJobInfoState{
			value: "SUCCEED",
		},
		FAILED: DigitalHumanBusinessCardJobInfoState{
			value: "FAILED",
		},
		CANCELED: DigitalHumanBusinessCardJobInfoState{
			value: "CANCELED",
		},
		BLOCK: DigitalHumanBusinessCardJobInfoState{
			value: "BLOCK",
		},
	}
}

func (c DigitalHumanBusinessCardJobInfoState) Value() string {
	return c.value
}

func (c DigitalHumanBusinessCardJobInfoState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DigitalHumanBusinessCardJobInfoState) UnmarshalJSON(b []byte) error {
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

type DigitalHumanBusinessCardJobInfoBusinessCardType struct {
	value string
}

type DigitalHumanBusinessCardJobInfoBusinessCardTypeEnum struct {
	E_2_D_DIGITAL_HUMAN_CARD DigitalHumanBusinessCardJobInfoBusinessCardType
}

func GetDigitalHumanBusinessCardJobInfoBusinessCardTypeEnum() DigitalHumanBusinessCardJobInfoBusinessCardTypeEnum {
	return DigitalHumanBusinessCardJobInfoBusinessCardTypeEnum{
		E_2_D_DIGITAL_HUMAN_CARD: DigitalHumanBusinessCardJobInfoBusinessCardType{
			value: "2D_DIGITAL_HUMAN_CARD",
		},
	}
}

func (c DigitalHumanBusinessCardJobInfoBusinessCardType) Value() string {
	return c.value
}

func (c DigitalHumanBusinessCardJobInfoBusinessCardType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DigitalHumanBusinessCardJobInfoBusinessCardType) UnmarshalJSON(b []byte) error {
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
