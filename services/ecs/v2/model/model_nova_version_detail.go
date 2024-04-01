package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// NovaVersionDetail
type NovaVersionDetail struct {

	// 所讨论的版本的通用名称。仅仅是信息性的，它没有真正的语义。
	Id string `json:"id"`

	// 链接到资源的问题。有关更多信息，请参见[OpenStack Documentation](http://developer.openstack.org/api-guide/compute/links_and_references.html)。
	Links []NovaLink `json:"links"`

	// 媒体类型。
	MediaTypes []NovaVersionMediaType `json:"media-types"`

	// 如果API的这个版本支持微版本，则支持最小的微版本。如果不支持微版本，这将是空字符串。
	MinVersion string `json:"min_version"`

	// 这个是API版本的状态。  可以是：  - CURRENT这是使用的API的首选版本； - SUPPORTED：这是一个较老的，但仍然支持的API版本； - DEPRECATED：一个被废弃的API版本，该版本将被删除
	Status NovaVersionDetailStatus `json:"status"`

	// 一个有特定值的字符串。API版本为2.0时，值为'2011-01-21T11:33:21Z' ，API版本是2.1时，值为' 2013-07-23T11:33:21Z'。
	Updated string `json:"updated"`

	// 如果API的这个版本支持微版本，则支持最大的微版本。如果不支持微版本，这将是空字符串。
	Version string `json:"version"`
}

func (o NovaVersionDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaVersionDetail struct{}"
	}

	return strings.Join([]string{"NovaVersionDetail", string(data)}, " ")
}

type NovaVersionDetailStatus struct {
	value string
}

type NovaVersionDetailStatusEnum struct {
	CURRENT    NovaVersionDetailStatus
	SUPPORTED  NovaVersionDetailStatus
	DEPRECATED NovaVersionDetailStatus
}

func GetNovaVersionDetailStatusEnum() NovaVersionDetailStatusEnum {
	return NovaVersionDetailStatusEnum{
		CURRENT: NovaVersionDetailStatus{
			value: "CURRENT",
		},
		SUPPORTED: NovaVersionDetailStatus{
			value: "SUPPORTED",
		},
		DEPRECATED: NovaVersionDetailStatus{
			value: "DEPRECATED",
		},
	}
}

func (c NovaVersionDetailStatus) Value() string {
	return c.value
}

func (c NovaVersionDetailStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NovaVersionDetailStatus) UnmarshalJSON(b []byte) error {
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
