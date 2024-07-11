package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AppDeleteResult 删除应用结果
type AppDeleteResult struct {

	// 应用id
	ApplicationId *string `json:"application_id,omitempty"`

	// 应用名称
	ApplicationName *string `json:"application_name,omitempty"`

	// 删除是否成功 success | error
	Status *AppDeleteResultStatus `json:"status,omitempty"`

	// 删除失败原因
	ErrorReason *string `json:"error_reason,omitempty"`
}

func (o AppDeleteResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppDeleteResult struct{}"
	}

	return strings.Join([]string{"AppDeleteResult", string(data)}, " ")
}

type AppDeleteResultStatus struct {
	value string
}

type AppDeleteResultStatusEnum struct {
	SUCCESS AppDeleteResultStatus
	ERROR   AppDeleteResultStatus
}

func GetAppDeleteResultStatusEnum() AppDeleteResultStatusEnum {
	return AppDeleteResultStatusEnum{
		SUCCESS: AppDeleteResultStatus{
			value: "success",
		},
		ERROR: AppDeleteResultStatus{
			value: "error",
		},
	}
}

func (c AppDeleteResultStatus) Value() string {
	return c.value
}

func (c AppDeleteResultStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AppDeleteResultStatus) UnmarshalJSON(b []byte) error {
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
