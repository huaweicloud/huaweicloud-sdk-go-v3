package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// UpdateDomainResponse Response Object
type UpdateDomainResponse struct {

	// 直播域名
	Domain *string `json:"domain,omitempty"`

	// 域名类型 - pull表示播放域名 - push表示推流域名
	DomainType *UpdateDomainResponseDomainType `json:"domain_type,omitempty"`

	// 直播域名的CNAME
	DomainCname *string `json:"domain_cname,omitempty"`

	// 直播所属直播中心
	Region *string `json:"region,omitempty"`

	// 直播域名的状态
	Status *UpdateDomainResponseStatus `json:"status,omitempty"`

	// 域名创建时间，格式：yyyy-mm-ddThh:mm:ssZ，UTC时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 状态描述
	StatusDescribe *string `json:"status_describe,omitempty"`

	// 域名应用区域 - mainland_china表示中国大陆区域 - outside_mainland_china表示中国大陆以外区域 - global表示全球加速区域
	ServiceArea *UpdateDomainResponseServiceArea `json:"service_area,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 域名支持的拉流协议；仅domain_type为pull时生效。若不填写此字段，视为默认支持FLV、RTMP拉流协议 - flv_rtmp表示支持FLV、RTMP协议 - hls表示支持HLS协议
	PullProtocol   *UpdateDomainResponsePullProtocol `json:"pull_protocol,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o UpdateDomainResponse) String() string {
	data, err := utils.Marshal(o)
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

func (c UpdateDomainResponseDomainType) Value() string {
	return c.value
}

func (c UpdateDomainResponseDomainType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateDomainResponseDomainType) UnmarshalJSON(b []byte) error {
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

type UpdateDomainResponseStatus struct {
	value string
}

type UpdateDomainResponseStatusEnum struct {
	ON          UpdateDomainResponseStatus
	OFF         UpdateDomainResponseStatus
	CONFIGURING UpdateDomainResponseStatus
	DISABLE     UpdateDomainResponseStatus
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
		DISABLE: UpdateDomainResponseStatus{
			value: "disable",
		},
	}
}

func (c UpdateDomainResponseStatus) Value() string {
	return c.value
}

func (c UpdateDomainResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateDomainResponseStatus) UnmarshalJSON(b []byte) error {
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

type UpdateDomainResponseServiceArea struct {
	value string
}

type UpdateDomainResponseServiceAreaEnum struct {
	MAINLAND_CHINA         UpdateDomainResponseServiceArea
	OUTSIDE_MAINLAND_CHINA UpdateDomainResponseServiceArea
	GLOBAL                 UpdateDomainResponseServiceArea
}

func GetUpdateDomainResponseServiceAreaEnum() UpdateDomainResponseServiceAreaEnum {
	return UpdateDomainResponseServiceAreaEnum{
		MAINLAND_CHINA: UpdateDomainResponseServiceArea{
			value: "mainland_china",
		},
		OUTSIDE_MAINLAND_CHINA: UpdateDomainResponseServiceArea{
			value: "outside_mainland_china",
		},
		GLOBAL: UpdateDomainResponseServiceArea{
			value: "global",
		},
	}
}

func (c UpdateDomainResponseServiceArea) Value() string {
	return c.value
}

func (c UpdateDomainResponseServiceArea) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateDomainResponseServiceArea) UnmarshalJSON(b []byte) error {
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

type UpdateDomainResponsePullProtocol struct {
	value string
}

type UpdateDomainResponsePullProtocolEnum struct {
	FLV_RTMP UpdateDomainResponsePullProtocol
	HLS      UpdateDomainResponsePullProtocol
}

func GetUpdateDomainResponsePullProtocolEnum() UpdateDomainResponsePullProtocolEnum {
	return UpdateDomainResponsePullProtocolEnum{
		FLV_RTMP: UpdateDomainResponsePullProtocol{
			value: "flv_rtmp",
		},
		HLS: UpdateDomainResponsePullProtocol{
			value: "hls",
		},
	}
}

func (c UpdateDomainResponsePullProtocol) Value() string {
	return c.value
}

func (c UpdateDomainResponsePullProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateDomainResponsePullProtocol) UnmarshalJSON(b []byte) error {
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
