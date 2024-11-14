package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListUsersOfStreamRequest Request Object
type ListUsersOfStreamRequest struct {

	// 播放域名。
	PlayDomain string `json:"play_domain"`

	// app名。
	App *string `json:"app,omitempty"`

	// 流名。
	Stream *string `json:"stream,omitempty"`

	// 运营商列表，取值如下： - CMCC ：移动 - CTCC ： 电信 - CUCC ：联通 - OTHER ：其他  不填写查询所有运营商。
	Isp *[]string `json:"isp,omitempty"`

	// 国家列表。具体取值请参考[国家名称缩写](https://support.huaweicloud.com/api-live/vod_08_0172.html)，不填写查询所有国家。
	Country *[]string `json:"country,omitempty"`

	// 区域列表。具体取值请参考[省份名称缩写](https://support.huaweicloud.com/api-live/live_03_0043.html)，不填写查询所有区域。
	Region *[]string `json:"region,omitempty"`

	// 请求协议
	Protocol *ListUsersOfStreamRequestProtocol `json:"protocol,omitempty"`

	// 查询数据的时间粒度，支持60（默认值）, 300秒。不传值时，使用默认值60秒。
	Interval *ListUsersOfStreamRequestInterval `json:"interval,omitempty"`

	// 起始时间。日期格式按照ISO8601表示法，并使用UTC时间。  格式为：YYYY-MM-DDThh:mm:ssZ。最大查询跨度31天，最大查询周期一年。  若参数为空，默认查询7天数据。
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间。日期格式按照ISO8601表示法，并使用UTC时间。格式为：YYYY-MM-DDThh:mm:ssZ。  若参数为空，默认为当前时间。结束时间需大于起始时间。
	EndTime *string `json:"end_time,omitempty"`

	// 服务类型： - Live：直播 - LLL：超低时延直播 - ALL: 默认所有直播
	ServiceType *ListUsersOfStreamRequestServiceType `json:"service_type,omitempty"`
}

func (o ListUsersOfStreamRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUsersOfStreamRequest struct{}"
	}

	return strings.Join([]string{"ListUsersOfStreamRequest", string(data)}, " ")
}

type ListUsersOfStreamRequestProtocol struct {
	value string
}

type ListUsersOfStreamRequestProtocolEnum struct {
	FLV ListUsersOfStreamRequestProtocol
	HLS ListUsersOfStreamRequestProtocol
}

func GetListUsersOfStreamRequestProtocolEnum() ListUsersOfStreamRequestProtocolEnum {
	return ListUsersOfStreamRequestProtocolEnum{
		FLV: ListUsersOfStreamRequestProtocol{
			value: "flv",
		},
		HLS: ListUsersOfStreamRequestProtocol{
			value: "hls",
		},
	}
}

func (c ListUsersOfStreamRequestProtocol) Value() string {
	return c.value
}

func (c ListUsersOfStreamRequestProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListUsersOfStreamRequestProtocol) UnmarshalJSON(b []byte) error {
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

type ListUsersOfStreamRequestInterval struct {
	value int32
}

type ListUsersOfStreamRequestIntervalEnum struct {
	E_60  ListUsersOfStreamRequestInterval
	E_300 ListUsersOfStreamRequestInterval
}

func GetListUsersOfStreamRequestIntervalEnum() ListUsersOfStreamRequestIntervalEnum {
	return ListUsersOfStreamRequestIntervalEnum{
		E_60: ListUsersOfStreamRequestInterval{
			value: 60,
		}, E_300: ListUsersOfStreamRequestInterval{
			value: 300,
		},
	}
}

func (c ListUsersOfStreamRequestInterval) Value() int32 {
	return c.value
}

func (c ListUsersOfStreamRequestInterval) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListUsersOfStreamRequestInterval) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type ListUsersOfStreamRequestServiceType struct {
	value string
}

type ListUsersOfStreamRequestServiceTypeEnum struct {
	LIVE ListUsersOfStreamRequestServiceType
	LLL  ListUsersOfStreamRequestServiceType
	ALL  ListUsersOfStreamRequestServiceType
}

func GetListUsersOfStreamRequestServiceTypeEnum() ListUsersOfStreamRequestServiceTypeEnum {
	return ListUsersOfStreamRequestServiceTypeEnum{
		LIVE: ListUsersOfStreamRequestServiceType{
			value: "Live",
		},
		LLL: ListUsersOfStreamRequestServiceType{
			value: "LLL",
		},
		ALL: ListUsersOfStreamRequestServiceType{
			value: "ALL",
		},
	}
}

func (c ListUsersOfStreamRequestServiceType) Value() string {
	return c.value
}

func (c ListUsersOfStreamRequestServiceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListUsersOfStreamRequestServiceType) UnmarshalJSON(b []byte) error {
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
