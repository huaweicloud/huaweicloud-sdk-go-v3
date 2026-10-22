package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AccessTypeStatistics 接入类型统计数据
type AccessTypeStatistics struct {

	// 接入类型（api/waf）
	AccessType AccessTypeStatisticsAccessType `json:"access_type"`

	// 用户输入请求数
	TextRequestCount int64 `json:"text_request_count"`

	// 用户输出请求数
	ImageRequestCount int64 `json:"image_request_count"`
}

func (o AccessTypeStatistics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessTypeStatistics struct{}"
	}

	return strings.Join([]string{"AccessTypeStatistics", string(data)}, " ")
}

type AccessTypeStatisticsAccessType struct {
	value string
}

type AccessTypeStatisticsAccessTypeEnum struct {
	API AccessTypeStatisticsAccessType
	WAF AccessTypeStatisticsAccessType
}

func GetAccessTypeStatisticsAccessTypeEnum() AccessTypeStatisticsAccessTypeEnum {
	return AccessTypeStatisticsAccessTypeEnum{
		API: AccessTypeStatisticsAccessType{
			value: "api",
		},
		WAF: AccessTypeStatisticsAccessType{
			value: "waf",
		},
	}
}

func (c AccessTypeStatisticsAccessType) Value() string {
	return c.value
}

func (c AccessTypeStatisticsAccessType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccessTypeStatisticsAccessType) UnmarshalJSON(b []byte) error {
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
