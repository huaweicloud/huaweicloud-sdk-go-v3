package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListTranscodeConcurrencyNumRequest Request Object
type ListTranscodeConcurrencyNumRequest struct {

	// 推流域名列表，最多支持查询100个域名，多个域名以逗号分隔。  若查询多个域名，则返回的是多个域名合并数据。
	PublishDomains []string `json:"publish_domains"`

	// 应用名称
	App *string `json:"app,omitempty"`

	// 查询数据的时间粒度。支持60, 300（默认值）和3600秒。不传值时，使用默认值300秒。
	Interval *ListTranscodeConcurrencyNumRequestInterval `json:"interval,omitempty"`

	// 起始时间。日期格式按照ISO8601表示法，并使用UTC时间。  格式为：YYYY-MM-DDThh:mm:ssZ。最大查询跨度1天，最大查询周期90天。  若参数为空，默认查询1天数据。
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间。日期格式按照ISO8601表示法，并使用UTC时间。  格式为：YYYY-MM-DDThh:mm:ssZ。最大查询跨度1天，最大查询周期90天。  若参数为空，默认为当前时间。结束时间需大于起始时间。
	EndTime *string `json:"end_time,omitempty"`
}

func (o ListTranscodeConcurrencyNumRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTranscodeConcurrencyNumRequest struct{}"
	}

	return strings.Join([]string{"ListTranscodeConcurrencyNumRequest", string(data)}, " ")
}

type ListTranscodeConcurrencyNumRequestInterval struct {
	value int32
}

type ListTranscodeConcurrencyNumRequestIntervalEnum struct {
	E_60   ListTranscodeConcurrencyNumRequestInterval
	E_300  ListTranscodeConcurrencyNumRequestInterval
	E_3600 ListTranscodeConcurrencyNumRequestInterval
}

func GetListTranscodeConcurrencyNumRequestIntervalEnum() ListTranscodeConcurrencyNumRequestIntervalEnum {
	return ListTranscodeConcurrencyNumRequestIntervalEnum{
		E_60: ListTranscodeConcurrencyNumRequestInterval{
			value: 60,
		}, E_300: ListTranscodeConcurrencyNumRequestInterval{
			value: 300,
		}, E_3600: ListTranscodeConcurrencyNumRequestInterval{
			value: 3600,
		},
	}
}

func (c ListTranscodeConcurrencyNumRequestInterval) Value() int32 {
	return c.value
}

func (c ListTranscodeConcurrencyNumRequestInterval) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTranscodeConcurrencyNumRequestInterval) UnmarshalJSON(b []byte) error {
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
