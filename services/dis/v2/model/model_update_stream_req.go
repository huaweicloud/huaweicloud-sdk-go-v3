package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateStreamReq struct {

	// 待更新的通道名称。
	StreamName string `json:"stream_name"`

	// 数据保留时长。  取值范围：24~72。  单位：小时。  空表示使用缺省值。
	DataDuration *int32 `json:"data_duration,omitempty"`

	// 源数据类型。  - BLOB：存储在数据库管理系统中的一组二进制数据。 - JSON：一种开放的文件格式，以易读的文字为基础，用来传输由属性值或者序列性的值组成的数据对象。 - CSV：纯文本形式存储的表格数据，分隔符默认采用逗号。  缺省值：BLOB。
	DataType *UpdateStreamReqDataType `json:"data_type,omitempty"`

	// 用于描述用户JSON、CSV格式的源数据结构，采用Avro Schema的语法描述。
	DataSchema *string `json:"data_schema,omitempty"`

	// 是否开启自动扩缩容。  - true：开启自动扩缩容。 - false：关闭自动扩缩容。  默认不开启。
	AutoScaleEnabled *bool `json:"auto_scale_enabled,omitempty"`

	// 当自动扩缩容启用时，自动缩容的最小分片数。
	AutoScaleMinPartitionCount *int64 `json:"auto_scale_min_partition_count,omitempty"`

	// 当自动扩缩容启用时，自动扩容的最大分片数。
	AutoScaleMaxPartitionCount *int64 `json:"auto_scale_max_partition_count,omitempty"`
}

func (o UpdateStreamReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStreamReq struct{}"
	}

	return strings.Join([]string{"UpdateStreamReq", string(data)}, " ")
}

type UpdateStreamReqDataType struct {
	value string
}

type UpdateStreamReqDataTypeEnum struct {
	BLOB UpdateStreamReqDataType
	JSON UpdateStreamReqDataType
	CSV  UpdateStreamReqDataType
}

func GetUpdateStreamReqDataTypeEnum() UpdateStreamReqDataTypeEnum {
	return UpdateStreamReqDataTypeEnum{
		BLOB: UpdateStreamReqDataType{
			value: "BLOB",
		},
		JSON: UpdateStreamReqDataType{
			value: "JSON",
		},
		CSV: UpdateStreamReqDataType{
			value: "CSV",
		},
	}
}

func (c UpdateStreamReqDataType) Value() string {
	return c.value
}

func (c UpdateStreamReqDataType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateStreamReqDataType) UnmarshalJSON(b []byte) error {
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
