package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type ListAreaDetailRequest struct {
	// 查询起始时间。日期格式按照ISO8601表示法，并使用UTC时间。  格式为：YYYY-MM-DDThh:mm:ssZ。最大查询跨度1天，最大查询周期90天。

	StartTime string `json:"start_time"`
	// 查询结束时间。日期格式按照ISO8601表示法，并使用UTC时间。  格式为：YYYY-MM-DDThh:mm:ssZ。最大查询跨度1天，最大查询周期90天。

	EndTime string `json:"end_time"`
	// 需查询的播放域名列表，最多支持查询10个域名。

	PlayDomains []string `json:"play_domains"`
	// 需查询的app。

	App *string `json:"app,omitempty"`
	// 流名称。

	Stream *string `json:"stream,omitempty"`
	// 查询数据的时间粒度。支持300（默认值）、3600和86400秒。若参数为空，则默认为300秒。  注意，若metric的值为player（观众数），则interval填入的值不起效果，查询粒度interval默认为60秒。

	Interval *ListAreaDetailRequestInterval `json:"interval,omitempty"`
	// 运营商列表，取值如下： - CMCC：移动 - CTCC：电信 - CUCC：联通 - OTHER：其他  若参数为空，则查询所有运营商。

	Isp *[]string `json:"isp,omitempty"`
	// 需查询的计费大区，取值如下： - CN - AP1 - AP2 - AP3 - EU - MEAA - NA - SA

	Area []string `json:"area"`
	// 指标，取值如下： - bandwidth：带宽 - traffic：流量 - player：观众数

	Metric string `json:"metric"`
	// 请求协议，取值如下： - flv - hls

	Protocol *string `json:"protocol,omitempty"`
}

func (o ListAreaDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAreaDetailRequest struct{}"
	}

	return strings.Join([]string{"ListAreaDetailRequest", string(data)}, " ")
}

type ListAreaDetailRequestInterval struct {
	value int32
}

type ListAreaDetailRequestIntervalEnum struct {
	E_300   ListAreaDetailRequestInterval
	E_3600  ListAreaDetailRequestInterval
	E_86400 ListAreaDetailRequestInterval
}

func GetListAreaDetailRequestIntervalEnum() ListAreaDetailRequestIntervalEnum {
	return ListAreaDetailRequestIntervalEnum{
		E_300: ListAreaDetailRequestInterval{
			value: 300,
		}, E_3600: ListAreaDetailRequestInterval{
			value: 3600,
		}, E_86400: ListAreaDetailRequestInterval{
			value: 86400,
		},
	}
}

func (c ListAreaDetailRequestInterval) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAreaDetailRequestInterval) UnmarshalJSON(b []byte) error {
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
