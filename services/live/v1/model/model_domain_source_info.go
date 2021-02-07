package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

type DomainSourceInfo struct {
	// 回源、转推协议。
	Protocol *DomainSourceInfoProtocol `json:"protocol,omitempty"`
	// 源站地址类型
	SourceType DomainSourceInfoSourceType `json:"source_type"`
	// 回源、转推地址列表，格式为：{domain/IP}[:{port}]，port默认值为1935；最少1个，最多10个。
	Sources *[]string `json:"sources,omitempty"`
	// 多个回源、转推地址的优先级。
	SourcesPriority *DomainSourceInfoSourcesPriority `json:"sources_priority,omitempty"`
}

func (o DomainSourceInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DomainSourceInfo struct{}"
	}

	return strings.Join([]string{"DomainSourceInfo", string(data)}, " ")
}

type DomainSourceInfoProtocol struct {
	value string
}

type DomainSourceInfoProtocolEnum struct {
	RTMP DomainSourceInfoProtocol
	HTTP DomainSourceInfoProtocol
}

func GetDomainSourceInfoProtocolEnum() DomainSourceInfoProtocolEnum {
	return DomainSourceInfoProtocolEnum{
		RTMP: DomainSourceInfoProtocol{
			value: "rtmp",
		},
		HTTP: DomainSourceInfoProtocol{
			value: "http",
		},
	}
}

func (c DomainSourceInfoProtocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *DomainSourceInfoProtocol) UnmarshalJSON(b []byte) error {
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

type DomainSourceInfoSourceType struct {
	value string
}

type DomainSourceInfoSourceTypeEnum struct {
	DOMAIN DomainSourceInfoSourceType
	IPADDR DomainSourceInfoSourceType
	HUAWEI DomainSourceInfoSourceType
}

func GetDomainSourceInfoSourceTypeEnum() DomainSourceInfoSourceTypeEnum {
	return DomainSourceInfoSourceTypeEnum{
		DOMAIN: DomainSourceInfoSourceType{
			value: "domain",
		},
		IPADDR: DomainSourceInfoSourceType{
			value: "ipaddr",
		},
		HUAWEI: DomainSourceInfoSourceType{
			value: "huawei",
		},
	}
}

func (c DomainSourceInfoSourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *DomainSourceInfoSourceType) UnmarshalJSON(b []byte) error {
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

type DomainSourceInfoSourcesPriority struct {
	value string
}

type DomainSourceInfoSourcesPriorityEnum struct {
	FIRST DomainSourceInfoSourcesPriority
	ALL   DomainSourceInfoSourcesPriority
}

func GetDomainSourceInfoSourcesPriorityEnum() DomainSourceInfoSourcesPriorityEnum {
	return DomainSourceInfoSourcesPriorityEnum{
		FIRST: DomainSourceInfoSourcesPriority{
			value: "first",
		},
		ALL: DomainSourceInfoSourcesPriority{
			value: "all",
		},
	}
}

func (c DomainSourceInfoSourcesPriority) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *DomainSourceInfoSourcesPriority) UnmarshalJSON(b []byte) error {
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
