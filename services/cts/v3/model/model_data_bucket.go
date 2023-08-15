package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type DataBucket struct {

	// 数据类追踪器追踪对象的桶名。 - 当启用或者停用数据类追踪器时，该参数为必选。 - 管理类追踪器无此参数。 - 追踪器一旦创建追踪桶无法修改。
	DataBucketName *string `json:"data_bucket_name,omitempty"`

	// 数据类追踪器追踪的操作类型。 - 当启用或者停用数据类追踪器时，该参数为必选。 - 管理类追踪器无此参数。
	DataEvent *[]DataBucketDataEvent `json:"data_event,omitempty"`
}

func (o DataBucket) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataBucket struct{}"
	}

	return strings.Join([]string{"DataBucket", string(data)}, " ")
}

type DataBucketDataEvent struct {
	value string
}

type DataBucketDataEventEnum struct {
	WRITE DataBucketDataEvent
	READ  DataBucketDataEvent
}

func GetDataBucketDataEventEnum() DataBucketDataEventEnum {
	return DataBucketDataEventEnum{
		WRITE: DataBucketDataEvent{
			value: "WRITE",
		},
		READ: DataBucketDataEvent{
			value: "READ",
		},
	}
}

func (c DataBucketDataEvent) Value() string {
	return c.value
}

func (c DataBucketDataEvent) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DataBucketDataEvent) UnmarshalJSON(b []byte) error {
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
