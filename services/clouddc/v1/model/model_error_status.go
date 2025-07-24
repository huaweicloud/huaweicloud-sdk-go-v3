package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ErrorStatus 失败时返回的错误对象  create_instance_error：创建实例失败   delete_instance_error：删除实例失败 reinstall_error：重新安装OS失败 modify_ip_error：修改ip失败  verify_server_error：验证服务器失败  delete_server_error：删除服务器失败
type ErrorStatus struct {

	// 错误码
	ErrorCode string `json:"error_code"`

	// 错误描述
	ErrorMsg string `json:"error_msg"`

	// 错误类型
	ErrorType ErrorStatusErrorType `json:"error_type"`
}

func (o ErrorStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ErrorStatus struct{}"
	}

	return strings.Join([]string{"ErrorStatus", string(data)}, " ")
}

type ErrorStatusErrorType struct {
	value string
}

type ErrorStatusErrorTypeEnum struct {
	CREATE_INSTANCE_ERROR ErrorStatusErrorType
	DELETE_INSTANCE_ERROR ErrorStatusErrorType
	REINSTALL_ERROR       ErrorStatusErrorType
	MODIFY_IP_ERROR       ErrorStatusErrorType
	VERIFY_SERVER_ERROR   ErrorStatusErrorType
	DELETE_SERVER_ERROR   ErrorStatusErrorType
}

func GetErrorStatusErrorTypeEnum() ErrorStatusErrorTypeEnum {
	return ErrorStatusErrorTypeEnum{
		CREATE_INSTANCE_ERROR: ErrorStatusErrorType{
			value: "create_instance_error",
		},
		DELETE_INSTANCE_ERROR: ErrorStatusErrorType{
			value: "delete_instance_error",
		},
		REINSTALL_ERROR: ErrorStatusErrorType{
			value: "reinstall_error",
		},
		MODIFY_IP_ERROR: ErrorStatusErrorType{
			value: "modify_ip_error",
		},
		VERIFY_SERVER_ERROR: ErrorStatusErrorType{
			value: "verify_server_error",
		},
		DELETE_SERVER_ERROR: ErrorStatusErrorType{
			value: "delete_server_error",
		},
	}
}

func (c ErrorStatusErrorType) Value() string {
	return c.value
}

func (c ErrorStatusErrorType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ErrorStatusErrorType) UnmarshalJSON(b []byte) error {
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
