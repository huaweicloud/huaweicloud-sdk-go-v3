package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

type LiveDomainCreateReq struct {
	// 直播域名
	Domain string `json:"domain"`
	// 域名类型 - pull表示播放域名 - push表示推流域名
	DomainType LiveDomainCreateReqDomainType `json:"domain_type"`
	// 直播所属的直播中心
	Region       string            `json:"region"`
	DomainSource *DomainSourceInfo `json:"domain_source,omitempty"`
}

func (o LiveDomainCreateReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "LiveDomainCreateReq struct{}"
	}

	return strings.Join([]string{"LiveDomainCreateReq", string(data)}, " ")
}

type LiveDomainCreateReqDomainType struct {
	value string
}

type LiveDomainCreateReqDomainTypeEnum struct {
	PULL LiveDomainCreateReqDomainType
	PUSH LiveDomainCreateReqDomainType
}

func GetLiveDomainCreateReqDomainTypeEnum() LiveDomainCreateReqDomainTypeEnum {
	return LiveDomainCreateReqDomainTypeEnum{
		PULL: LiveDomainCreateReqDomainType{
			value: "pull",
		},
		PUSH: LiveDomainCreateReqDomainType{
			value: "push",
		},
	}
}

func (c LiveDomainCreateReqDomainType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *LiveDomainCreateReqDomainType) UnmarshalJSON(b []byte) error {
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
