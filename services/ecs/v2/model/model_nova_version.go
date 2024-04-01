package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// NovaVersion
type NovaVersion struct {

	// 所讨论的版本的通用名称。仅仅是信息性的，它没有真正的语义。
	Id string `json:"id"`

	// 版本相关标记快捷链接信息。
	Links []NovaLink `json:"links"`

	// 如果API的这个版本支持微版本，则支持最小的微版本。如果不支持微版本，这将是空字符串。
	MinVersion string `json:"min_version"`

	// 这个是API版本的状态。可以是：  - CURRENT，这是使用的API的首选版本 - SUPPORTED，这是一个较老的，但仍然支持的API版本 - DEPRECATED，一个被废弃的API版本，该版本将被删除
	Status NovaVersionStatus `json:"status"`

	// 如果API的这个版本支持微版本，则支持最大的微版本。如果不支持微版本，这将是空字符串。
	Version string `json:"version"`

	// 一个有特定值的字符串。API版本为2.0时，值为'2011-01-21T11:33:21Z'，API版本是2.1时，值为' 2013-07-23T11:33:21Z'。
	Updated string `json:"updated"`
}

func (o NovaVersion) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaVersion struct{}"
	}

	return strings.Join([]string{"NovaVersion", string(data)}, " ")
}

type NovaVersionStatus struct {
	value string
}

type NovaVersionStatusEnum struct {
	CURRENT    NovaVersionStatus
	SUPPORTED  NovaVersionStatus
	DEPRECATED NovaVersionStatus
}

func GetNovaVersionStatusEnum() NovaVersionStatusEnum {
	return NovaVersionStatusEnum{
		CURRENT: NovaVersionStatus{
			value: "CURRENT",
		},
		SUPPORTED: NovaVersionStatus{
			value: "SUPPORTED",
		},
		DEPRECATED: NovaVersionStatus{
			value: "DEPRECATED",
		},
	}
}

func (c NovaVersionStatus) Value() string {
	return c.value
}

func (c NovaVersionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NovaVersionStatus) UnmarshalJSON(b []byte) error {
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
