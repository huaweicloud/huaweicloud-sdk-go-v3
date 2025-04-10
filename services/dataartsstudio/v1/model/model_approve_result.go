package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ApproveResult 审批结果
type ApproveResult struct {

	// 工单id
	Id *string `json:"id,omitempty"`

	// 审批结果,SUCCESS,FAIL
	Status *ApproveResultStatus `json:"status,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`
}

func (o ApproveResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApproveResult struct{}"
	}

	return strings.Join([]string{"ApproveResult", string(data)}, " ")
}

type ApproveResultStatus struct {
	value string
}

type ApproveResultStatusEnum struct {
	SUCCESS ApproveResultStatus
	FAIL    ApproveResultStatus
}

func GetApproveResultStatusEnum() ApproveResultStatusEnum {
	return ApproveResultStatusEnum{
		SUCCESS: ApproveResultStatus{
			value: "SUCCESS",
		},
		FAIL: ApproveResultStatus{
			value: "FAIL",
		},
	}
}

func (c ApproveResultStatus) Value() string {
	return c.value
}

func (c ApproveResultStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApproveResultStatus) UnmarshalJSON(b []byte) error {
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
