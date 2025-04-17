package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type RefreshResult struct {

	// 刷新URL。
	Url *string `json:"url,omitempty"`

	// 刷新任务状态。  取值如下： - PROCESSING：处理中 - SUCCEED：刷新完成 - FAILED：刷新失败
	Status *RefreshResultStatus `json:"status,omitempty"`
}

func (o RefreshResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RefreshResult struct{}"
	}

	return strings.Join([]string{"RefreshResult", string(data)}, " ")
}

type RefreshResultStatus struct {
	value string
}

type RefreshResultStatusEnum struct {
	PROCESSING RefreshResultStatus
	SUCCEED    RefreshResultStatus
	FAILED     RefreshResultStatus
}

func GetRefreshResultStatusEnum() RefreshResultStatusEnum {
	return RefreshResultStatusEnum{
		PROCESSING: RefreshResultStatus{
			value: "PROCESSING",
		},
		SUCCEED: RefreshResultStatus{
			value: "SUCCEED",
		},
		FAILED: RefreshResultStatus{
			value: "FAILED",
		},
	}
}

func (c RefreshResultStatus) Value() string {
	return c.value
}

func (c RefreshResultStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RefreshResultStatus) UnmarshalJSON(b []byte) error {
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
