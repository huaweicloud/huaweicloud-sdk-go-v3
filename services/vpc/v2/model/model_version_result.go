package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// VersionResult
type VersionResult struct {

	// API版本的状态：  - CURRENT（当前版本）  - STABLE（稳定版本）  - DEPRECATED（废弃版本）
	Status VersionResultStatus `json:"status"`

	// API版本
	Id string `json:"id"`

	// 链接列表
	Links []NeutronPageLink `json:"links"`
}

func (o VersionResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionResult struct{}"
	}

	return strings.Join([]string{"VersionResult", string(data)}, " ")
}

type VersionResultStatus struct {
	value string
}

type VersionResultStatusEnum struct {
	CURRENT    VersionResultStatus
	STABLE     VersionResultStatus
	DEPRECATED VersionResultStatus
}

func GetVersionResultStatusEnum() VersionResultStatusEnum {
	return VersionResultStatusEnum{
		CURRENT: VersionResultStatus{
			value: "CURRENT",
		},
		STABLE: VersionResultStatus{
			value: "STABLE",
		},
		DEPRECATED: VersionResultStatus{
			value: "DEPRECATED",
		},
	}
}

func (c VersionResultStatus) Value() string {
	return c.value
}

func (c VersionResultStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VersionResultStatus) UnmarshalJSON(b []byte) error {
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
