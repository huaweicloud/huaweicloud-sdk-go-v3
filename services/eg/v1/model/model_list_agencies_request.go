package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListAgenciesRequest struct {

	// 服务委托授权场景类型
	Type ListAgenciesRequestType `json:"type"`
}

func (o ListAgenciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAgenciesRequest struct{}"
	}

	return strings.Join([]string{"ListAgenciesRequest", string(data)}, " ")
}

type ListAgenciesRequestType struct {
	value string
}

type ListAgenciesRequestTypeEnum struct {
	TARGET_CONNECTION      ListAgenciesRequestType
	CUSTOM_SOURCE_RABBITMQ ListAgenciesRequestType
	EG_RESTORE_AGENCY      ListAgenciesRequestType
}

func GetListAgenciesRequestTypeEnum() ListAgenciesRequestTypeEnum {
	return ListAgenciesRequestTypeEnum{
		TARGET_CONNECTION: ListAgenciesRequestType{
			value: "TARGET_CONNECTION",
		},
		CUSTOM_SOURCE_RABBITMQ: ListAgenciesRequestType{
			value: "CUSTOM_SOURCE_RABBITMQ",
		},
		EG_RESTORE_AGENCY: ListAgenciesRequestType{
			value: "EG_RESTORE_AGENCY",
		},
	}
}

func (c ListAgenciesRequestType) Value() string {
	return c.value
}

func (c ListAgenciesRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAgenciesRequestType) UnmarshalJSON(b []byte) error {
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
