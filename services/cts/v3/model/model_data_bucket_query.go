package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 追踪桶信息。
type DataBucketQuery struct {
	// 标识OBS桶名称。由数字或字母开头，支持小写字母、数字、“-”、“.”，长度为3～63个字符。

	DataBucketName *string `json:"data_bucket_name,omitempty"`
	// 追踪桶日志是否支持搜索。

	SearchEnabled *bool `json:"search_enabled,omitempty"`
	// 数据类追踪器追踪对象的桶名。 - 当启用或者停用数据类追踪器时，该参数为必选。 - 管理类追踪器无此参数。 - 追踪器一旦创建追踪桶无法修改。

	DataEvent *[]DataBucketQueryDataEvent `json:"data_event,omitempty"`
}

func (o DataBucketQuery) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataBucketQuery struct{}"
	}

	return strings.Join([]string{"DataBucketQuery", string(data)}, " ")
}

type DataBucketQueryDataEvent struct {
	value string
}

type DataBucketQueryDataEventEnum struct {
	WRITE DataBucketQueryDataEvent
	READ  DataBucketQueryDataEvent
}

func GetDataBucketQueryDataEventEnum() DataBucketQueryDataEventEnum {
	return DataBucketQueryDataEventEnum{
		WRITE: DataBucketQueryDataEvent{
			value: "WRITE",
		},
		READ: DataBucketQueryDataEvent{
			value: "READ",
		},
	}
}

func (c DataBucketQueryDataEvent) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DataBucketQueryDataEvent) UnmarshalJSON(b []byte) error {
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
