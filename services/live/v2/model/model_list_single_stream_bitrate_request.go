package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSingleStreamBitrateRequest Request Object
type ListSingleStreamBitrateRequest struct {

	// 推流域名。
	Domain string `json:"domain"`

	// App名。
	App string `json:"app"`

	// 流名。
	Stream string `json:"stream"`

	// 数据类型，取值如下：  - VIDEO ：视频码率  - AUDIO ：音频码率   不填写默认查询视频码率的数据。
	Type *ListSingleStreamBitrateRequestType `json:"type,omitempty"`

	// 起始时间。日期格式按照ISO8601表示法，并使用UTC时间。  格式为：YYYY-MM-DDThh:mm:ssZ。最大查询跨度1天，最大查询周期1个月。  若参数为空，默认查询最近1小时数据。
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间。日期格式按照ISO8601表示法，并使用UTC时间。格式为：YYYY-MM-DDThh:mm:ssZ。  若参数为空，默认为当前时间。结束时间需大于起始时间。
	EndTime *string `json:"end_time,omitempty"`
}

func (o ListSingleStreamBitrateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSingleStreamBitrateRequest struct{}"
	}

	return strings.Join([]string{"ListSingleStreamBitrateRequest", string(data)}, " ")
}

type ListSingleStreamBitrateRequestType struct {
	value string
}

type ListSingleStreamBitrateRequestTypeEnum struct {
	VIDEO ListSingleStreamBitrateRequestType
	AUDIO ListSingleStreamBitrateRequestType
}

func GetListSingleStreamBitrateRequestTypeEnum() ListSingleStreamBitrateRequestTypeEnum {
	return ListSingleStreamBitrateRequestTypeEnum{
		VIDEO: ListSingleStreamBitrateRequestType{
			value: "VIDEO",
		},
		AUDIO: ListSingleStreamBitrateRequestType{
			value: "AUDIO",
		},
	}
}

func (c ListSingleStreamBitrateRequestType) Value() string {
	return c.value
}

func (c ListSingleStreamBitrateRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSingleStreamBitrateRequestType) UnmarshalJSON(b []byte) error {
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
