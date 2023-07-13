package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListStorageTypesRequest Request Object
type ListStorageTypesRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	// 数据库版本号。
	Version string `json:"version"`

	// 实例类型： enterprise(企业版)， centralization_standard(主备版)，不区分大小写。
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
	ENTERPRISE              ListStorageTypesRequestHaMode
	CENTRALIZATION_STANDARD ListStorageTypesRequestHaMode
}

func GetListStorageTypesRequestHaModeEnum() ListStorageTypesRequestHaModeEnum {
	return ListStorageTypesRequestHaModeEnum{
		ENTERPRISE: ListStorageTypesRequestHaMode{
			value: "enterprise",
		},
		CENTRALIZATION_STANDARD: ListStorageTypesRequestHaMode{
			value: "centralization_standard",
		},
	}
}

func (c ListStorageTypesRequestHaMode) Value() string {
	return c.value
}

func (c ListStorageTypesRequestHaMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListStorageTypesRequestHaMode) UnmarshalJSON(b []byte) error {
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
