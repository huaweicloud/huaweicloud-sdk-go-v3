package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListDomainBandwidthPeakRequest Request Object
type ListDomainBandwidthPeakRequest struct {

	// 播放域名列表，最多支持查询100个域名，多个域名以逗号分隔。  如果不传入域名，则查询租户下所有播放域名的带宽峰值。
	PlayDomains *[]string `json:"play_domains,omitempty"`

	// 应用名称。
	App *string `json:"app,omitempty"`

	// 流名。
	Stream *string `json:"stream,omitempty"`

	// 区域列表。具体取值请参考[省份名称缩写](live_03_0043.xml)，不填写查询所有区域。
	Region *[]string `json:"region,omitempty"`

	// 运营商列表，取值如下： - CMCC ：移动 - CTCC ： 电信 - CUCC ：联通 - OTHER ：其他  不填写查询所有运营商。
	Isp *[]string `json:"isp,omitempty"`

	// 请求协议
	Protocol *ListDomainBandwidthPeakRequestProtocol `json:"protocol,omitempty"`

	// 起始时间。日期格式按照ISO8601表示法，并使用UTC时间。  格式为：YYYY-MM-DDThh:mm:ssZ。最大查询跨度31天，最大查询周期一年。  若参数为空，默认查询7天数据。
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间。日期格式按照ISO8601表示法，并使用UTC时间。格式为：YYYY-MM-DDThh:mm:ssZ。  若参数为空，默认为当前时间。结束时间需大于起始时间。
	EndTime *string `json:"end_time,omitempty"`

	// 服务类型： - Live：直播 - LLL：超低时延直播 - ALL: 默认所有直播
	ServiceType *ListDomainBandwidthPeakRequestServiceType `json:"service_type,omitempty"`
}

func (o ListDomainBandwidthPeakRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainBandwidthPeakRequest struct{}"
	}

	return strings.Join([]string{"ListDomainBandwidthPeakRequest", string(data)}, " ")
}

type ListDomainBandwidthPeakRequestProtocol struct {
	value string
}

type ListDomainBandwidthPeakRequestProtocolEnum struct {
	FLV ListDomainBandwidthPeakRequestProtocol
	HLS ListDomainBandwidthPeakRequestProtocol
}

func GetListDomainBandwidthPeakRequestProtocolEnum() ListDomainBandwidthPeakRequestProtocolEnum {
	return ListDomainBandwidthPeakRequestProtocolEnum{
		FLV: ListDomainBandwidthPeakRequestProtocol{
			value: "flv",
		},
		HLS: ListDomainBandwidthPeakRequestProtocol{
			value: "hls",
		},
	}
}

func (c ListDomainBandwidthPeakRequestProtocol) Value() string {
	return c.value
}

func (c ListDomainBandwidthPeakRequestProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDomainBandwidthPeakRequestProtocol) UnmarshalJSON(b []byte) error {
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

type ListDomainBandwidthPeakRequestServiceType struct {
	value string
}

type ListDomainBandwidthPeakRequestServiceTypeEnum struct {
	LIVE ListDomainBandwidthPeakRequestServiceType
	LLL  ListDomainBandwidthPeakRequestServiceType
	ALL  ListDomainBandwidthPeakRequestServiceType
}

func GetListDomainBandwidthPeakRequestServiceTypeEnum() ListDomainBandwidthPeakRequestServiceTypeEnum {
	return ListDomainBandwidthPeakRequestServiceTypeEnum{
		LIVE: ListDomainBandwidthPeakRequestServiceType{
			value: "Live",
		},
		LLL: ListDomainBandwidthPeakRequestServiceType{
			value: "LLL",
		},
		ALL: ListDomainBandwidthPeakRequestServiceType{
			value: "ALL",
		},
	}
}

func (c ListDomainBandwidthPeakRequestServiceType) Value() string {
	return c.value
}

func (c ListDomainBandwidthPeakRequestServiceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDomainBandwidthPeakRequestServiceType) UnmarshalJSON(b []byte) error {
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
