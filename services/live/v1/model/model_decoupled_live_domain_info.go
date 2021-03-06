package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"strings"
)

type DecoupledLiveDomainInfo struct {
	// 直播域名

	Domain string `json:"domain"`
	// 域名类型

	DomainType DecoupledLiveDomainInfoDomainType `json:"domain_type"`
	// CDN厂商

	Vendor DecoupledLiveDomainInfoVendor `json:"vendor"`
	// 直播所属直播中心

	Region string `json:"region"`
	// 直播域名的CName

	DomainCname string `json:"domain_cname"`
	// 直播域名的状态

	Status DecoupledLiveDomainInfoStatus `json:"status"`
	// 播放域名关联的推流域名（只有domain_type为pull的时候有效）

	RelatedDomain string `json:"related_domain"`
	// 域名创建时间，格式：yyyy-mm-ddThh:mm:ssZ，UTC时间

	CreateTime *sdktime.SdkTime `json:"create_time"`

	DomainSource *DomainSourceInfo `json:"domain_source"`
}

func (o DecoupledLiveDomainInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DecoupledLiveDomainInfo struct{}"
	}

	return strings.Join([]string{"DecoupledLiveDomainInfo", string(data)}, " ")
}

type DecoupledLiveDomainInfoDomainType struct {
	value string
}

type DecoupledLiveDomainInfoDomainTypeEnum struct {
	PULL DecoupledLiveDomainInfoDomainType
	PUSH DecoupledLiveDomainInfoDomainType
}

func GetDecoupledLiveDomainInfoDomainTypeEnum() DecoupledLiveDomainInfoDomainTypeEnum {
	return DecoupledLiveDomainInfoDomainTypeEnum{
		PULL: DecoupledLiveDomainInfoDomainType{
			value: "pull",
		},
		PUSH: DecoupledLiveDomainInfoDomainType{
			value: "push",
		},
	}
}

func (c DecoupledLiveDomainInfoDomainType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *DecoupledLiveDomainInfoDomainType) UnmarshalJSON(b []byte) error {
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

type DecoupledLiveDomainInfoVendor struct {
	value string
}

type DecoupledLiveDomainInfoVendorEnum struct {
	CHINA_NET_CENTER DecoupledLiveDomainInfoVendor
	KING_SOFT        DecoupledLiveDomainInfoVendor
	TENCENT          DecoupledLiveDomainInfoVendor
	U_CDN            DecoupledLiveDomainInfoVendor
	TENCENT_OUTSIDE  DecoupledLiveDomainInfoVendor
	DNION            DecoupledLiveDomainInfoVendor
	BAISHAN          DecoupledLiveDomainInfoVendor
	BAIDU            DecoupledLiveDomainInfoVendor
	ONETHING         DecoupledLiveDomainInfoVendor
}

func GetDecoupledLiveDomainInfoVendorEnum() DecoupledLiveDomainInfoVendorEnum {
	return DecoupledLiveDomainInfoVendorEnum{
		CHINA_NET_CENTER: DecoupledLiveDomainInfoVendor{
			value: "ChinaNetCenter",
		},
		KING_SOFT: DecoupledLiveDomainInfoVendor{
			value: "KingSoft",
		},
		TENCENT: DecoupledLiveDomainInfoVendor{
			value: "Tencent",
		},
		U_CDN: DecoupledLiveDomainInfoVendor{
			value: "uCDN",
		},
		TENCENT_OUTSIDE: DecoupledLiveDomainInfoVendor{
			value: "TencentOutside",
		},
		DNION: DecoupledLiveDomainInfoVendor{
			value: "Dnion",
		},
		BAISHAN: DecoupledLiveDomainInfoVendor{
			value: "Baishan",
		},
		BAIDU: DecoupledLiveDomainInfoVendor{
			value: "Baidu",
		},
		ONETHING: DecoupledLiveDomainInfoVendor{
			value: "Onething",
		},
	}
}

func (c DecoupledLiveDomainInfoVendor) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *DecoupledLiveDomainInfoVendor) UnmarshalJSON(b []byte) error {
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

type DecoupledLiveDomainInfoStatus struct {
	value string
}

type DecoupledLiveDomainInfoStatusEnum struct {
	ON          DecoupledLiveDomainInfoStatus
	OFF         DecoupledLiveDomainInfoStatus
	CONFIGURING DecoupledLiveDomainInfoStatus
}

func GetDecoupledLiveDomainInfoStatusEnum() DecoupledLiveDomainInfoStatusEnum {
	return DecoupledLiveDomainInfoStatusEnum{
		ON: DecoupledLiveDomainInfoStatus{
			value: "on",
		},
		OFF: DecoupledLiveDomainInfoStatus{
			value: "off",
		},
		CONFIGURING: DecoupledLiveDomainInfoStatus{
			value: "configuring",
		},
	}
}

func (c DecoupledLiveDomainInfoStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *DecoupledLiveDomainInfoStatus) UnmarshalJSON(b []byte) error {
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
