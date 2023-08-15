package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// QueryNetworkResult 测试连接结果响应体。
type QueryNetworkResult struct {

	// 测试连接IP。
	Ip *string `json:"ip,omitempty"`

	// 测试连接是否成功。
	Success bool `json:"success"`

	// 测试连接结果。
	Result *string `json:"result,omitempty"`

	// 测试连接是否成功。取值： - success：成功。 - failed：不成功。
	Status QueryNetworkResultStatus `json:"status"`

	// 测试连接失败错误码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 测试连接失败错误内容。
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o QueryNetworkResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryNetworkResult struct{}"
	}

	return strings.Join([]string{"QueryNetworkResult", string(data)}, " ")
}

type QueryNetworkResultStatus struct {
	value string
}

type QueryNetworkResultStatusEnum struct {
	SUCCESS QueryNetworkResultStatus
	FAILED  QueryNetworkResultStatus
}

func GetQueryNetworkResultStatusEnum() QueryNetworkResultStatusEnum {
	return QueryNetworkResultStatusEnum{
		SUCCESS: QueryNetworkResultStatus{
			value: "success",
		},
		FAILED: QueryNetworkResultStatus{
			value: "failed",
		},
	}
}

func (c QueryNetworkResultStatus) Value() string {
	return c.value
}

func (c QueryNetworkResultStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryNetworkResultStatus) UnmarshalJSON(b []byte) error {
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
