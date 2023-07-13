package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AgencyCreateReq 服务委托创建请求
type AgencyCreateReq struct {

	// 服务委托应用场景类型
	Type AgencyCreateReqType `json:"type"`
}

func (o AgencyCreateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgencyCreateReq struct{}"
	}

	return strings.Join([]string{"AgencyCreateReq", string(data)}, " ")
}

type AgencyCreateReqType struct {
	value string
}

type AgencyCreateReqTypeEnum struct {
	TARGET_CONNECTION      AgencyCreateReqType
	CUSTOM_SOURCE_RABBITMQ AgencyCreateReqType
	EG_RESTORE_AGENCY      AgencyCreateReqType
}

func GetAgencyCreateReqTypeEnum() AgencyCreateReqTypeEnum {
	return AgencyCreateReqTypeEnum{
		TARGET_CONNECTION: AgencyCreateReqType{
			value: "TARGET_CONNECTION",
		},
		CUSTOM_SOURCE_RABBITMQ: AgencyCreateReqType{
			value: "CUSTOM_SOURCE_RABBITMQ",
		},
		EG_RESTORE_AGENCY: AgencyCreateReqType{
			value: "EG_RESTORE_AGENCY",
		},
	}
}

func (c AgencyCreateReqType) Value() string {
	return c.value
}

func (c AgencyCreateReqType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AgencyCreateReqType) UnmarshalJSON(b []byte) error {
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
