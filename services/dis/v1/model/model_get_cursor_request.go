package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type GetCursorRequest struct {
	// 已创建的通道名称。

	StreamName string `json:"stream-name"`
	// 通道的分区标识符。  可定义为如下两种样式：  - shardId-0000000000 - 0  比如一个通道有三个分区，那么分区标识符分别为0, 1, 2，或者shardId-0000000000, shardId-0000000001, shardId-0000000002

	PartitionId string `json:"partition-id"`
	// 游标类型。  - AT_SEQUENCE_NUMBER：从特定序列号（即starting-sequence-number定义的序列号）所在的记录开始读取数据。此类型为默认游标类型。  - AFTER_SEQUENCE_NUMBER：从特定序列号（即starting-sequence-number定义的序列号）后的记录开始读取数据。  - TRIM_HORIZON：从最早被存储至分区的有效记录开始读取。例如，某租户使用DIS的通道，分别上传了三条数据A1，A2，A3。N天后（设定A1已过期，A2和A3仍在有效期范围内），该租户需要下载此三条数据，并选择了TRIM_HORIZON这种下载方式。那么用户可下载的数据将从A2开始读取。  - LATEST：从分区中的最新记录开始读取，此设置可以保证你总是读到分区中最新记录。  - AT_TIMESTAMP：从特定时间戳（即timestamp定义的时间戳）开始读取。

	CursorType *GetCursorRequestCursorType `json:"cursor-type,omitempty"`
	// 序列号。序列号是每个记录的唯一标识符。序列号由DIS在数据生产者调用PutRecords操作以添加数据到DIS数据通道时DIS服务自动分配的。同一分区键的序列号通常会随时间变化增加。PutRecords请求之间的时间段越长，序列号越大。  序列号与游标类型AT_SEQUENCE_NUMBER和AFTER_SEQUENCE_NUMBER强相关，二者共同确定读取数据的位置。  取值范围：0~9223372036854775807。

	StartingSequenceNumber *string `json:"starting-sequence-number,omitempty"`
	// 开始读取数据记录的时间戳，与游标类型AT_TIMESTAMP强相关，二者共同确定读取数据的位置。  说明：  此时间戳精确到毫秒。

	Timestamp *int64 `json:"timestamp,omitempty"`
}

func (o GetCursorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetCursorRequest struct{}"
	}

	return strings.Join([]string{"GetCursorRequest", string(data)}, " ")
}

type GetCursorRequestCursorType struct {
	value string
}

type GetCursorRequestCursorTypeEnum struct {
	AT_SEQUENCE_NUMBER    GetCursorRequestCursorType
	AFTER_SEQUENCE_NUMBER GetCursorRequestCursorType
	TRIM_HORIZON          GetCursorRequestCursorType
	LATEST                GetCursorRequestCursorType
	AT_TIMESTAMP          GetCursorRequestCursorType
}

func GetGetCursorRequestCursorTypeEnum() GetCursorRequestCursorTypeEnum {
	return GetCursorRequestCursorTypeEnum{
		AT_SEQUENCE_NUMBER: GetCursorRequestCursorType{
			value: "AT_SEQUENCE_NUMBER",
		},
		AFTER_SEQUENCE_NUMBER: GetCursorRequestCursorType{
			value: "AFTER_SEQUENCE_NUMBER",
		},
		TRIM_HORIZON: GetCursorRequestCursorType{
			value: "TRIM_HORIZON",
		},
		LATEST: GetCursorRequestCursorType{
			value: "LATEST",
		},
		AT_TIMESTAMP: GetCursorRequestCursorType{
			value: "AT_TIMESTAMP",
		},
	}
}

func (c GetCursorRequestCursorType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GetCursorRequestCursorType) UnmarshalJSON(b []byte) error {
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
