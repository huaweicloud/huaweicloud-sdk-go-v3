package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 正则化参数
type Regular struct {

	// L2正则项系数。
	L2Regularization *float64 `json:"l2_regularization,omitempty"`

	// 正则损失计算方式。
	RegularLossComputeMode *RegularRegularLossComputeMode `json:"regular_loss_compute_mode,omitempty"`

	// 隐向量层L2正则化系数。
	EmbedL2Regularization *float64 `json:"embed_l2_regularization,omitempty"`

	// wide部分L2正则化系数。
	WideL2Regularization *float64 `json:"wide_l2_regularization,omitempty"`

	// 结构化部分L2正则化系数。
	StructureL2Regularization *float64 `json:"structure_l2_regularization,omitempty"`
}

func (o Regular) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Regular struct{}"
	}

	return strings.Join([]string{"Regular", string(data)}, " ")
}

type RegularRegularLossComputeMode struct {
	value string
}

type RegularRegularLossComputeModeEnum struct {
	FULL  RegularRegularLossComputeMode
	BATCH RegularRegularLossComputeMode
}

func GetRegularRegularLossComputeModeEnum() RegularRegularLossComputeModeEnum {
	return RegularRegularLossComputeModeEnum{
		FULL: RegularRegularLossComputeMode{
			value: "full",
		},
		BATCH: RegularRegularLossComputeMode{
			value: "batch",
		},
	}
}

func (c RegularRegularLossComputeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RegularRegularLossComputeMode) UnmarshalJSON(b []byte) error {
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
