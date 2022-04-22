package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListDomainBandwidthPeakRequest struct {

	// 播放域名列表，最多支持查询100个域名，多个域名以逗号分隔。
	PlayDomains []string `json:"play_domains"`

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

func (c ListDomainBandwidthPeakRequestProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDomainBandwidthPeakRequestProtocol) UnmarshalJSON(b []byte) error {
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
