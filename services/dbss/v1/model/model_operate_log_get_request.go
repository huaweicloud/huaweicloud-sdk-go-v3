package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type OperateLogGetRequest struct {
	Time *TimeRangeBean `json:"time"`

	// 操作日志用户名
	UserName *string `json:"user_name,omitempty"`

	// 动作名称 - CREATE - DELETE - DOWNLOAD - UPDATE
	Action *string `json:"action,omitempty"`

	// 执行结果 - success - fail
	Result *OperateLogGetRequestResult `json:"result,omitempty"`

	// 页数
	Page string `json:"page"`

	// 每页条数
	Size string `json:"size"`
}

func (o OperateLogGetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperateLogGetRequest struct{}"
	}

	return strings.Join([]string{"OperateLogGetRequest", string(data)}, " ")
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
