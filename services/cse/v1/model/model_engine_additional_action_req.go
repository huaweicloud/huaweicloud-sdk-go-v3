package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// EngineAdditionalActionReq 升级微服务引擎请求体
type EngineAdditionalActionReq struct {

	// 操作类型
	Action EngineAdditionalActionReqAction `json:"action"`
}

func (o EngineAdditionalActionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EngineAdditionalActionReq struct{}"
	}

	return strings.Join([]string{"EngineAdditionalActionReq", string(data)}, " ")
}

type EngineAdditionalActionReqAction struct {
	value string
}

type EngineAdditionalActionReqActionEnum struct {
	RETRY EngineAdditionalActionReqAction
}

func GetEngineAdditionalActionReqActionEnum() EngineAdditionalActionReqActionEnum {
	return EngineAdditionalActionReqActionEnum{
		RETRY: EngineAdditionalActionReqAction{
			value: "Retry",
		},
	}
}

func (c EngineAdditionalActionReqAction) Value() string {
	return c.value
}

func (c EngineAdditionalActionReqAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EngineAdditionalActionReqAction) UnmarshalJSON(b []byte) error {
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
