package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowStreamMetricsRequest Request Object
type ShowStreamMetricsRequest struct {

	// 通道名称。
	StreamName string `json:"stream_name"`

	// 通道监控指标。（label与label_list必须二选一，label_list与label同时存在时，以label_list为准）  - total_put_bytes_per_stream：总输入流量（Byte） - total_get_bytes_per_stream：总输出流量（Byte） - total_put_records_per_stream：总输入记录数（个） - total_get_records_per_stream：总输出记录数（个） - total_put_req_latency：上传请求平均处理时间（毫秒） - total_get_req_latency：下载请求平均处理时间（毫秒） - total_put_req_suc_per_stream：上传请求成功次数（个） - total_get_req_suc_per_stream：下载请求成功次数（个） - traffic_control_put：因流控拒绝的上传请求次数 （个） - traffic_control_get：因流控拒绝的下载请求次数 （个）
	Label *ShowStreamMetricsRequestLabel `json:"label,omitempty"`

	// 使用label用逗号拼接组成，用于批量查询多个label的指标。（label与label_list必须二选一，label_list与label同时存在时，以label_list为准）
	LabelList *string `json:"label_list,omitempty"`

	// 监控开始时间点，10位时间戳。
	StartTime int64 `json:"start_time"`

	// 监控结束时间点，10位时间戳。
	EndTime int64 `json:"end_time"`
}

func (o ShowStreamMetricsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStreamMetricsRequest struct{}"
	}

	return strings.Join([]string{"ShowStreamMetricsRequest", string(data)}, " ")
}

type ShowStreamMetricsRequestLabel struct {
	value string
}

type ShowStreamMetricsRequestLabelEnum struct {
	TOTAL_PUT_BYTES_PER_STREAM   ShowStreamMetricsRequestLabel
	TOTAL_GET_BYTES_PER_STREAM   ShowStreamMetricsRequestLabel
	TOTAL_PUT_RECORDS_PER_STREAM ShowStreamMetricsRequestLabel
	TOTAL_GET_RECORDS_PER_STREAM ShowStreamMetricsRequestLabel
	TOTAL_PUT_REQ_LATENCY        ShowStreamMetricsRequestLabel
	TOTAL_GET_REQ_LATENCY        ShowStreamMetricsRequestLabel
	TOTAL_PUT_REQ_SUC_PER_STREAM ShowStreamMetricsRequestLabel
	TOTAL_GET_REQ_SUC_PER_STREAM ShowStreamMetricsRequestLabel
	TRAFFIC_CONTROL_PUT          ShowStreamMetricsRequestLabel
	TRAFFIC_CONTROL_GET          ShowStreamMetricsRequestLabel
}

func GetShowStreamMetricsRequestLabelEnum() ShowStreamMetricsRequestLabelEnum {
	return ShowStreamMetricsRequestLabelEnum{
		TOTAL_PUT_BYTES_PER_STREAM: ShowStreamMetricsRequestLabel{
			value: "total_put_bytes_per_stream",
		},
		TOTAL_GET_BYTES_PER_STREAM: ShowStreamMetricsRequestLabel{
			value: "total_get_bytes_per_stream",
		},
		TOTAL_PUT_RECORDS_PER_STREAM: ShowStreamMetricsRequestLabel{
			value: "total_put_records_per_stream",
		},
		TOTAL_GET_RECORDS_PER_STREAM: ShowStreamMetricsRequestLabel{
			value: "total_get_records_per_stream",
		},
		TOTAL_PUT_REQ_LATENCY: ShowStreamMetricsRequestLabel{
			value: "total_put_req_latency",
		},
		TOTAL_GET_REQ_LATENCY: ShowStreamMetricsRequestLabel{
			value: "total_get_req_latency",
		},
		TOTAL_PUT_REQ_SUC_PER_STREAM: ShowStreamMetricsRequestLabel{
			value: "total_put_req_suc_per_stream",
		},
		TOTAL_GET_REQ_SUC_PER_STREAM: ShowStreamMetricsRequestLabel{
			value: "total_get_req_suc_per_stream",
		},
		TRAFFIC_CONTROL_PUT: ShowStreamMetricsRequestLabel{
			value: "traffic_control_put",
		},
		TRAFFIC_CONTROL_GET: ShowStreamMetricsRequestLabel{
			value: "traffic_control_get",
		},
	}
}

func (c ShowStreamMetricsRequestLabel) Value() string {
	return c.value
}

func (c ShowStreamMetricsRequestLabel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowStreamMetricsRequestLabel) UnmarshalJSON(b []byte) error {
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
