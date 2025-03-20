package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

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
	Region string `json:"region"`

	// 域名应用区域 - mainland_china表示中国大陆区域 - outside_mainland_china表示中国大陆以外区域 - global表示全球加速区域
	ServiceArea *LiveDomainCreateReqServiceArea `json:"service_area,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 域名支持的拉流协议；仅domain_type为pull时生效。若不填写此字段，视为默认支持FLV、RTMP拉流协议 - flv_rtmp表示支持FLV、RTMP协议 - hls表示支持HLS协议
	PullProtocol *LiveDomainCreateReqPullProtocol `json:"pull_protocol,omitempty"`
}

func (o LiveDomainCreateReq) String() string {
	data, err := utils.Marshal(o)
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

func (c LiveDomainCreateReqDomainType) Value() string {
	return c.value
}

func (c LiveDomainCreateReqDomainType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LiveDomainCreateReqDomainType) UnmarshalJSON(b []byte) error {
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

type LiveDomainCreateReqServiceArea struct {
	value string
}

type LiveDomainCreateReqServiceAreaEnum struct {
	MAINLAND_CHINA         LiveDomainCreateReqServiceArea
	OUTSIDE_MAINLAND_CHINA LiveDomainCreateReqServiceArea
	GLOBAL                 LiveDomainCreateReqServiceArea
}

func GetLiveDomainCreateReqServiceAreaEnum() LiveDomainCreateReqServiceAreaEnum {
	return LiveDomainCreateReqServiceAreaEnum{
		MAINLAND_CHINA: LiveDomainCreateReqServiceArea{
			value: "mainland_china",
		},
		OUTSIDE_MAINLAND_CHINA: LiveDomainCreateReqServiceArea{
			value: "outside_mainland_china",
		},
		GLOBAL: LiveDomainCreateReqServiceArea{
			value: "global",
		},
	}
}

func (c LiveDomainCreateReqServiceArea) Value() string {
	return c.value
}

func (c LiveDomainCreateReqServiceArea) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LiveDomainCreateReqServiceArea) UnmarshalJSON(b []byte) error {
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

type LiveDomainCreateReqPullProtocol struct {
	value string
}

type LiveDomainCreateReqPullProtocolEnum struct {
	FLV_RTMP LiveDomainCreateReqPullProtocol
	HLS      LiveDomainCreateReqPullProtocol
}

func GetLiveDomainCreateReqPullProtocolEnum() LiveDomainCreateReqPullProtocolEnum {
	return LiveDomainCreateReqPullProtocolEnum{
		FLV_RTMP: LiveDomainCreateReqPullProtocol{
			value: "flv_rtmp",
		},
		HLS: LiveDomainCreateReqPullProtocol{
			value: "hls",
		},
	}
}

func (c LiveDomainCreateReqPullProtocol) Value() string {
	return c.value
}

func (c LiveDomainCreateReqPullProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LiveDomainCreateReqPullProtocol) UnmarshalJSON(b []byte) error {
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
