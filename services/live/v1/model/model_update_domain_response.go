package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"strings"
)

// Response Object
type UpdateDomainResponse struct {
	// 直播域名
	Domain *string `json:"domain,omitempty"`
	// 域名类型 - pull表示播放域名 - push表示推流域名
	DomainType *UpdateDomainResponseDomainType `json:"domain_type,omitempty"`
	// 直播域名的CName
	DomainCname *string `json:"domain_cname,omitempty"`
	// 直播所属直播中心
	Region *string `json:"region,omitempty"`
	// 直播域名的状态
	Status *UpdateDomainResponseStatus `json:"status,omitempty"`
	// 域名创建时间，格式：yyyy-mm-ddThh:mm:ssZ，UTC时间
	CreateTime     *sdktime.SdkTime  `json:"create_time,omitempty"`
	DomainSource   *DomainSourceInfo `json:"domain_source,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o UpdateDomainResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateDomainResponse struct{}"
	}

	return strings.Join([]string{"UpdateDomainResponse", string(data)}, " ")
}

type UpdateDomainResponseDomainType struct {
	value string
}

type UpdateDomainResponseDomainTypeEnum struct {
	PULL UpdateDomainResponseDomainType
	PUSH UpdateDomainResponseDomainType
}

func GetUpdateDomainResponseDomainTypeEnum() UpdateDomainResponseDomainTypeEnum {
	return UpdateDomainResponseDomainTypeEnum{
		PULL: UpdateDomainResponseDomainType{
			value: "pull",
		},
		PUSH: UpdateDomainResponseDomainType{
			value: "push",
		},
	}
}

func (c UpdateDomainResponseDomainType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *UpdateDomainResponseDomainType) UnmarshalJSON(b []byte) error {
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

type UpdateDomainResponseStatus struct {
	value string
}

type UpdateDomainResponseStatusEnum struct {
	ON          UpdateDomainResponseStatus
	OFF         UpdateDomainResponseStatus
	CONFIGURING UpdateDomainResponseStatus
}

func GetUpdateDomainResponseStatusEnum() UpdateDomainResponseStatusEnum {
	return UpdateDomainResponseStatusEnum{
		ON: UpdateDomainResponseStatus{
			value: "on",
		},
		OFF: UpdateDomainResponseStatus{
			value: "off",
		},
		CONFIGURING: UpdateDomainResponseStatus{
			value: "configuring",
		},
	}
}

func (c UpdateDomainResponseStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *UpdateDomainResponseStatus) UnmarshalJSON(b []byte) error {
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
