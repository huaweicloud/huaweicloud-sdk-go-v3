package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateConnectivityTestResponse Response Object
type CreateConnectivityTestResponse struct {

	// 创建备份导入页面实例连接测试结果
	Result         *CreateConnectivityTestResponseResult `json:"result,omitempty"`
	HttpStatusCode int                                   `json:"-"`
}

func (o CreateConnectivityTestResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConnectivityTestResponse struct{}"
	}

	return strings.Join([]string{"CreateConnectivityTestResponse", string(data)}, " ")
}

type CreateConnectivityTestResponseResult struct {
	value string
}

type CreateConnectivityTestResponseResultEnum struct {
	SUCCESS CreateConnectivityTestResponseResult
	FAILED  CreateConnectivityTestResponseResult
}

func GetCreateConnectivityTestResponseResultEnum() CreateConnectivityTestResponseResultEnum {
	return CreateConnectivityTestResponseResultEnum{
		SUCCESS: CreateConnectivityTestResponseResult{
			value: "SUCCESS",
		},
		FAILED: CreateConnectivityTestResponseResult{
			value: "FAILED",
		},
	}
}

func (c CreateConnectivityTestResponseResult) Value() string {
	return c.value
}

func (c CreateConnectivityTestResponseResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateConnectivityTestResponseResult) UnmarshalJSON(b []byte) error {
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
