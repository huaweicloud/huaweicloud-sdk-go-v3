package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type DomainItem struct {

	// 高危漏洞数
	High *int32 `json:"high,omitempty"`

	// 中危漏洞数
	Middle *int32 `json:"middle,omitempty"`

	// 低危漏洞数
	Low *int32 `json:"low,omitempty"`

	// 提示危漏洞数
	Hint *int32 `json:"hint,omitempty"`

	// 域名id
	DomainId *string `json:"domain_id,omitempty"`

	// 一级域名id
	TopLevelDomainId *string `json:"top_level_domain_id,omitempty"`

	// 域名
	DomainName *string `json:"domain_name,omitempty"`

	// 域名的别名
	Alias *string `json:"alias,omitempty"`

	// 创建域名资产的时间
	CreateTime *string `json:"create_time,omitempty"`

	// 域名的认证状态:   * unauth - 未认证   * auth - 已认证   * invalid - 认证文件无效   * manual - 人工认证   * skip - 免认证
	AuthStatus *DomainItemAuthStatus `json:"auth_status,omitempty"`

	// 协议类型:   * http:// - HTTP   * https:// - HTTPS
	ProtocolType *DomainItemProtocolType `json:"protocol_type,omitempty"`
}

func (o DomainItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainItem struct{}"
	}

	return strings.Join([]string{"DomainItem", string(data)}, " ")
}

type DomainItemAuthStatus struct {
	value string
}

type DomainItemAuthStatusEnum struct {
	UNAUTH  DomainItemAuthStatus
	AUTH    DomainItemAuthStatus
	INVALID DomainItemAuthStatus
	MANUAL  DomainItemAuthStatus
	SKIP    DomainItemAuthStatus
}

func GetDomainItemAuthStatusEnum() DomainItemAuthStatusEnum {
	return DomainItemAuthStatusEnum{
		UNAUTH: DomainItemAuthStatus{
			value: "unauth",
		},
		AUTH: DomainItemAuthStatus{
			value: "auth",
		},
		INVALID: DomainItemAuthStatus{
			value: "invalid",
		},
		MANUAL: DomainItemAuthStatus{
			value: "manual",
		},
		SKIP: DomainItemAuthStatus{
			value: "skip",
		},
	}
}

func (c DomainItemAuthStatus) Value() string {
	return c.value
}

func (c DomainItemAuthStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DomainItemAuthStatus) UnmarshalJSON(b []byte) error {
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

type DomainItemProtocolType struct {
	value string
}

type DomainItemProtocolTypeEnum struct {
	HTTP_  DomainItemProtocolType
	HTTPS_ DomainItemProtocolType
}

func GetDomainItemProtocolTypeEnum() DomainItemProtocolTypeEnum {
	return DomainItemProtocolTypeEnum{
		HTTP_: DomainItemProtocolType{
			value: "http://",
		},
		HTTPS_: DomainItemProtocolType{
			value: "https://",
		},
	}
}

func (c DomainItemProtocolType) Value() string {
	return c.value
}

func (c DomainItemProtocolType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DomainItemProtocolType) UnmarshalJSON(b []byte) error {
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
