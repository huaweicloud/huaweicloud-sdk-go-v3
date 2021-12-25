package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type CancelSqlResponse struct {
	// 错误信息。

	Message *string `json:"message,omitempty"`
	// 取消SQL的执行结果。  说明： 默认返回SUCCEED，对于已经结束的任务也会返回SUCCEED，只有取消正在运行的SQL时没成功才会FAILED。  枚举值： - SUCCEED：成功 - FAILED：失败

	Status         *CancelSqlResponseStatus `json:"status,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o CancelSqlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelSqlResponse struct{}"
	}

	return strings.Join([]string{"CancelSqlResponse", string(data)}, " ")
}

type CancelSqlResponseStatus struct {
	value string
}

type CancelSqlResponseStatusEnum struct {
	SUCCEED CancelSqlResponseStatus
	FAILED  CancelSqlResponseStatus
}

func GetCancelSqlResponseStatusEnum() CancelSqlResponseStatusEnum {
	return CancelSqlResponseStatusEnum{
		SUCCEED: CancelSqlResponseStatus{
			value: "SUCCEED",
		},
		FAILED: CancelSqlResponseStatus{
			value: "FAILED",
		},
	}
}

func (c CancelSqlResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CancelSqlResponseStatus) UnmarshalJSON(b []byte) error {
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
