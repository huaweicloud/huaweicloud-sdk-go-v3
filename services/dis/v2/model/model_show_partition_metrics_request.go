package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowPartitionMetricsRequest Request Object
type ShowPartitionMetricsRequest struct {

	// 通道名称。
	StreamName string `json:"stream_name"`

	// 分区编号。 可定义为如下两种样式： - shardId-0000000000 - 0 比如一个通道有三个分区，那么分区标识符分别为0, 1, 2，或者shardId-0000000000, shardId-0000000001, shardId-0000000002
	PartitionId string `json:"partition_id"`

	// 分区监控指标。（label与label_list必须二选一，label_list与label同时存在时，以label_list为准）  - total_put_bytes_per_partition：分区总输入流量（Byte） - total_get_bytes_per_partition：分区总输出流量（Byte） - total_put_records_per_partition：分区总输入记录数（个） - total_get_records_per_partition：分区总输出记录数（个）
	Label *ShowPartitionMetricsRequestLabel `json:"label,omitempty"`

	// 使用label用逗号拼接组成，用于批量查询多个label的指标。（label与label_list必须二选一，label_list与label同时存在时，以label_list为准）
	LabelList *string `json:"label_list,omitempty"`

	// 监控开始时间点，10位时间戳。
	StartTime int64 `json:"start_time"`

	// 监控结束时间点，10位时间戳。
	EndTime string `json:"end_time"`
}

func (o ShowPartitionMetricsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPartitionMetricsRequest struct{}"
	}

	return strings.Join([]string{"ShowPartitionMetricsRequest", string(data)}, " ")
}

type ShowPartitionMetricsRequestLabel struct {
	value string
}

type ShowPartitionMetricsRequestLabelEnum struct {
	TOTAL_PUT_BYTES_PER_PARTITION   ShowPartitionMetricsRequestLabel
	TOTAL_GET_BYTES_PER_PARTITION   ShowPartitionMetricsRequestLabel
	TOTAL_PUT_RECORDS_PER_PARTITION ShowPartitionMetricsRequestLabel
	TOTAL_GET_RECORDS_PER_PARTITION ShowPartitionMetricsRequestLabel
}

func GetShowPartitionMetricsRequestLabelEnum() ShowPartitionMetricsRequestLabelEnum {
	return ShowPartitionMetricsRequestLabelEnum{
		TOTAL_PUT_BYTES_PER_PARTITION: ShowPartitionMetricsRequestLabel{
			value: "total_put_bytes_per_partition",
		},
		TOTAL_GET_BYTES_PER_PARTITION: ShowPartitionMetricsRequestLabel{
			value: "total_get_bytes_per_partition",
		},
		TOTAL_PUT_RECORDS_PER_PARTITION: ShowPartitionMetricsRequestLabel{
			value: "total_put_records_per_partition",
		},
		TOTAL_GET_RECORDS_PER_PARTITION: ShowPartitionMetricsRequestLabel{
			value: "total_get_records_per_partition",
		},
	}
}

func (c ShowPartitionMetricsRequestLabel) Value() string {
	return c.value
}

func (c ShowPartitionMetricsRequestLabel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowPartitionMetricsRequestLabel) UnmarshalJSON(b []byte) error {
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
