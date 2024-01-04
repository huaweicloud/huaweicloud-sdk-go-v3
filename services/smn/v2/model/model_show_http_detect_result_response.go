package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowHttpDetectResultResponse Response Object
type ShowHttpDetectResultResponse struct {
	Detail *GetHttpDetectResponseBodyDetail `json:"detail,omitempty"`

	// http探测任务状态，0代表执行成功，终端可用，1代表未执行，2代表执行失败，终端不可用
	Status *ShowHttpDetectResultResponseStatus `json:"status,omitempty"`

	// 请求ID
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowHttpDetectResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttpDetectResultResponse struct{}"
	}

	return strings.Join([]string{"ShowHttpDetectResultResponse", string(data)}, " ")
}

type ShowHttpDetectResultResponseStatus struct {
	value int32
}

type ShowHttpDetectResultResponseStatusEnum struct {
	E_0 ShowHttpDetectResultResponseStatus
	E_1 ShowHttpDetectResultResponseStatus
	E_2 ShowHttpDetectResultResponseStatus
}

func GetShowHttpDetectResultResponseStatusEnum() ShowHttpDetectResultResponseStatusEnum {
	return ShowHttpDetectResultResponseStatusEnum{
		E_0: ShowHttpDetectResultResponseStatus{
			value: 0,
		}, E_1: ShowHttpDetectResultResponseStatus{
			value: 1,
		}, E_2: ShowHttpDetectResultResponseStatus{
			value: 2,
		},
	}
}

func (c ShowHttpDetectResultResponseStatus) Value() int32 {
	return c.value
}

func (c ShowHttpDetectResultResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowHttpDetectResultResponseStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
