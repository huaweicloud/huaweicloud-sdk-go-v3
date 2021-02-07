package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"strings"
)

// Response Object
type CreateDomainResponse struct {
	// 直播域名
	Domain *string `json:"domain,omitempty"`
	// 域名类型 - pull表示播放域名 - push表示推流域名
	DomainType *CreateDomainResponseDomainType `json:"domain_type,omitempty"`
	// 直播域名的CName
	DomainCname *string `json:"domain_cname,omitempty"`
	// 直播所属直播中心
	Region *string `json:"region,omitempty"`
	// 直播域名的状态
	Status *CreateDomainResponseStatus `json:"status,omitempty"`
	// 域名创建时间，格式：yyyy-mm-ddThh:mm:ssZ，UTC时间
	CreateTime     *sdktime.SdkTime  `json:"create_time,omitempty"`
	DomainSource   *DomainSourceInfo `json:"domain_source,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o CreateDomainResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateDomainResponse struct{}"
	}

	return strings.Join([]string{"CreateDomainResponse", string(data)}, " ")
}

type CreateDomainResponseDomainType struct {
	value string
}

type CreateDomainResponseDomainTypeEnum struct {
	PULL CreateDomainResponseDomainType
	PUSH CreateDomainResponseDomainType
}

func GetCreateDomainResponseDomainTypeEnum() CreateDomainResponseDomainTypeEnum {
	return CreateDomainResponseDomainTypeEnum{
		PULL: CreateDomainResponseDomainType{
			value: "pull",
		},
		PUSH: CreateDomainResponseDomainType{
			value: "push",
		},
	}
}

func (c CreateDomainResponseDomainType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *CreateDomainResponseDomainType) UnmarshalJSON(b []byte) error {
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

type CreateDomainResponseStatus struct {
	value string
}

type CreateDomainResponseStatusEnum struct {
	ON          CreateDomainResponseStatus
	OFF         CreateDomainResponseStatus
	CONFIGURING CreateDomainResponseStatus
}

func GetCreateDomainResponseStatusEnum() CreateDomainResponseStatusEnum {
	return CreateDomainResponseStatusEnum{
		ON: CreateDomainResponseStatus{
			value: "on",
		},
		OFF: CreateDomainResponseStatus{
			value: "off",
		},
		CONFIGURING: CreateDomainResponseStatus{
			value: "configuring",
		},
	}
}

func (c CreateDomainResponseStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *CreateDomainResponseStatus) UnmarshalJSON(b []byte) error {
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
