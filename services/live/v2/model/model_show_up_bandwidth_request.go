package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type ShowUpBandwidthRequest struct {
	// 推流域名列表，最多支持查询100个域名，多个域名以逗号分隔，若查询多个域名，则返回的是多个域名合并数据。

	PublishDomains []string `json:"publish_domains"`
	// 应用名称。

	App *string `json:"app,omitempty"`
	// 流名。

	Stream *string `json:"stream,omitempty"`
	// 区域列表。具体取值请参考[省份名称缩写](live_03_0043.xml)，不填写查询所有区域。

	Region *[]string `json:"region,omitempty"`
	// 运营商列表，取值如下： - CMCC ：移动 - CTCC ： 电信 - CUCC ：联通 - OTHER ：其他  不填写查询所有运营商。

	Isp *[]string `json:"isp,omitempty"`
	// 查询数据的时间粒度。支持300（默认值），3600和86400秒。不传值时，使用默认值300秒。

	Interval *ShowUpBandwidthRequestInterval `json:"interval,omitempty"`
	// 起始时间。日期格式按照ISO8601表示法，并使用UTC时间。  格式为：YYYY-MM-DDThh:mm:ssZ。 最大查询跨度31天，最大查询周期1年。  若参数为空，默认查询7天数据。

	StartTime *string `json:"start_time,omitempty"`
	// 结束时间。日期格式按照ISO8601表示法，并使用UTC时间。 格式为：YYYY-MM-DDThh:mm:ssZ。  若参数为空，默认为当前时间，最大查询跨度31天，最大查询周期1年。结束时间需大于起始时间。

	EndTime *string `json:"end_time,omitempty"`
}

func (o ShowUpBandwidthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUpBandwidthRequest struct{}"
	}

	return strings.Join([]string{"ShowUpBandwidthRequest", string(data)}, " ")
}

type ShowUpBandwidthRequestInterval struct {
	value int32
}

type ShowUpBandwidthRequestIntervalEnum struct {
	E_300   ShowUpBandwidthRequestInterval
	E_3600  ShowUpBandwidthRequestInterval
	E_86400 ShowUpBandwidthRequestInterval
}

func GetShowUpBandwidthRequestIntervalEnum() ShowUpBandwidthRequestIntervalEnum {
	return ShowUpBandwidthRequestIntervalEnum{
		E_300: ShowUpBandwidthRequestInterval{
			value: 300,
		}, E_3600: ShowUpBandwidthRequestInterval{
			value: 3600,
		}, E_86400: ShowUpBandwidthRequestInterval{
			value: 86400,
		},
	}
}

func (c ShowUpBandwidthRequestInterval) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowUpBandwidthRequestInterval) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
