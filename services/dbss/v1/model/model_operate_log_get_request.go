package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type OperateLogGetRequest struct {
	Time *TimeRangeBean `json:"time,omitempty"`

	// 操作日志用户名
	UserName *string `json:"user_name,omitempty"`

	// 动作名称 - CREATE - DELETE - DOWNLOAD - UPDATE
	Action *OperateLogGetRequestAction `json:"action,omitempty"`

	// 执行结果 - success - fail
	Result *OperateLogGetRequestResult `json:"result,omitempty"`

	// 页数
	Page *string `json:"page,omitempty"`

	// 每页条数
	Size *string `json:"size,omitempty"`
}

func (o OperateLogGetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperateLogGetRequest struct{}"
	}

	return strings.Join([]string{"OperateLogGetRequest", string(data)}, " ")
}

type OperateLogGetRequestAction struct {
	value string
}

type OperateLogGetRequestActionEnum struct {
	CREATE   OperateLogGetRequestAction
	DELETE   OperateLogGetRequestAction
	DOWNLOAD OperateLogGetRequestAction
	UPDATE   OperateLogGetRequestAction
}

func GetOperateLogGetRequestActionEnum() OperateLogGetRequestActionEnum {
	return OperateLogGetRequestActionEnum{
		CREATE: OperateLogGetRequestAction{
			value: "CREATE",
		},
		DELETE: OperateLogGetRequestAction{
			value: "DELETE",
		},
		DOWNLOAD: OperateLogGetRequestAction{
			value: "DOWNLOAD",
		},
		UPDATE: OperateLogGetRequestAction{
			value: "UPDATE",
		},
	}
}

func (c OperateLogGetRequestAction) Value() string {
	return c.value
}

func (c OperateLogGetRequestAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OperateLogGetRequestAction) UnmarshalJSON(b []byte) error {
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

type OperateLogGetRequestResult struct {
	value string
}

type OperateLogGetRequestResultEnum struct {
	SUCCESS OperateLogGetRequestResult
	FAIL    OperateLogGetRequestResult
}

func GetOperateLogGetRequestResultEnum() OperateLogGetRequestResultEnum {
	return OperateLogGetRequestResultEnum{
		SUCCESS: OperateLogGetRequestResult{
			value: "success",
		},
		FAIL: OperateLogGetRequestResult{
			value: "fail",
		},
	}
}

func (c OperateLogGetRequestResult) Value() string {
	return c.value
}

func (c OperateLogGetRequestResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OperateLogGetRequestResult) UnmarshalJSON(b []byte) error {
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
