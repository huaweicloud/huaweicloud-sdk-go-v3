package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type ListStorageTypesRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`
	// 数据库版本号。

	Version string `json:"version"`
	// 实例类型： enterprise(企业版)， ha(主备版)，不区分大小写。

	HaMode *ListStorageTypesRequestHaMode `json:"ha_mode,omitempty"`
}

func (o ListStorageTypesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStorageTypesRequest struct{}"
	}

	return strings.Join([]string{"ListStorageTypesRequest", string(data)}, " ")
}

type ListStorageTypesRequestHaMode struct {
	value string
}

type ListStorageTypesRequestHaModeEnum struct {
	ENTERPRISE ListStorageTypesRequestHaMode
	HA         ListStorageTypesRequestHaMode
}

func GetListStorageTypesRequestHaModeEnum() ListStorageTypesRequestHaModeEnum {
	return ListStorageTypesRequestHaModeEnum{
		ENTERPRISE: ListStorageTypesRequestHaMode{
			value: "enterprise",
		},
		HA: ListStorageTypesRequestHaMode{
			value: "ha",
		},
	}
}

func (c ListStorageTypesRequestHaMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListStorageTypesRequestHaMode) UnmarshalJSON(b []byte) error {
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
