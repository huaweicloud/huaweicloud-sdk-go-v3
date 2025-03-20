package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
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

	// 状态描述
	StatusDescribe *string `json:"status_describe,omitempty"`

	// 域名应用区域 - mainland_china表示中国大陆区域 - outside_mainland_china表示中国大陆以外区域 - global表示全球加速区域
	ServiceArea *DecoupledLiveDomainInfoServiceArea `json:"service_area,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 域名支持的拉流协议；仅domain_type为pull时生效。若不填写此字段，视为默认支持FLV、RTMP拉流协议 - flv_rtmp表示支持FLV、RTMP协议 - hls表示支持HLS协议
	PullProtocol *DecoupledLiveDomainInfoPullProtocol `json:"pull_protocol,omitempty"`

	// IPV6开关是否开启，默认关闭，true为开启；false或空为关闭
	IsIpv6 *bool `json:"is_ipv6,omitempty"`
}

func (o DecoupledLiveDomainInfo) String() string {
	data, err := utils.Marshal(o)
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

func (c DecoupledLiveDomainInfoDomainType) Value() string {
	return c.value
}

func (c DecoupledLiveDomainInfoDomainType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DecoupledLiveDomainInfoDomainType) UnmarshalJSON(b []byte) error {
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

func (c DecoupledLiveDomainInfoVendor) Value() string {
	return c.value
}

func (c DecoupledLiveDomainInfoVendor) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DecoupledLiveDomainInfoVendor) UnmarshalJSON(b []byte) error {
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

func (c DecoupledLiveDomainInfoStatus) Value() string {
	return c.value
}

func (c DecoupledLiveDomainInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DecoupledLiveDomainInfoStatus) UnmarshalJSON(b []byte) error {
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

type DecoupledLiveDomainInfoServiceArea struct {
	value string
}

type DecoupledLiveDomainInfoServiceAreaEnum struct {
	MAINLAND_CHINA         DecoupledLiveDomainInfoServiceArea
	OUTSIDE_MAINLAND_CHINA DecoupledLiveDomainInfoServiceArea
	GLOBAL                 DecoupledLiveDomainInfoServiceArea
}

func GetDecoupledLiveDomainInfoServiceAreaEnum() DecoupledLiveDomainInfoServiceAreaEnum {
	return DecoupledLiveDomainInfoServiceAreaEnum{
		MAINLAND_CHINA: DecoupledLiveDomainInfoServiceArea{
			value: "mainland_china",
		},
		OUTSIDE_MAINLAND_CHINA: DecoupledLiveDomainInfoServiceArea{
			value: "outside_mainland_china",
		},
		GLOBAL: DecoupledLiveDomainInfoServiceArea{
			value: "global",
		},
	}
}

func (c DecoupledLiveDomainInfoServiceArea) Value() string {
	return c.value
}

func (c DecoupledLiveDomainInfoServiceArea) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DecoupledLiveDomainInfoServiceArea) UnmarshalJSON(b []byte) error {
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

type DecoupledLiveDomainInfoPullProtocol struct {
	value string
}

type DecoupledLiveDomainInfoPullProtocolEnum struct {
	FLV_RTMP DecoupledLiveDomainInfoPullProtocol
	HLS      DecoupledLiveDomainInfoPullProtocol
}

func GetDecoupledLiveDomainInfoPullProtocolEnum() DecoupledLiveDomainInfoPullProtocolEnum {
	return DecoupledLiveDomainInfoPullProtocolEnum{
		FLV_RTMP: DecoupledLiveDomainInfoPullProtocol{
			value: "flv_rtmp",
		},
		HLS: DecoupledLiveDomainInfoPullProtocol{
			value: "hls",
		},
	}
}

func (c DecoupledLiveDomainInfoPullProtocol) Value() string {
	return c.value
}

func (c DecoupledLiveDomainInfoPullProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DecoupledLiveDomainInfoPullProtocol) UnmarshalJSON(b []byte) error {
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
