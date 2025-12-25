package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type MetricMetaDataDerivedMetrics struct {

	// 衍生指标结果维度，0维：单个数字，2维：图表或表格，3+维：多标签图表
	MetricDimension int32 `json:"metric_dimension"`

	// 指标支持的最大检索范围，单位：天；
	MaxQueryRange *int32 `json:"max_query_range,omitempty"`

	// 指标查询范围相对起始时间 datemath表达式
	DateStart *string `json:"date_start,omitempty"`

	// 指标查询范围相对截止时间 datemath表达式
	DateEnd *string `json:"date_end,omitempty"`

	// 时间格式，epoch_millis;epoch_second;yyyy-MM-dd'T'HH:mm:ss.SSSZ
	DateFormat *string `json:"date_format,omitempty"`

	// 获取指标结果方式，cbsl, api, dsl, sql
	QueryType MetricMetaDataDerivedMetricsQueryType `json:"query_type"`

	// 获取指标结果的方法，转义成字符串, 共四种query方式：CBSL、API、DSL、SQL - query_type为`CBSL`时，function传入dataspace_id，pipe_id，query，sort, from, to   样例：     ```     {\\\"dataspace_id\\\":\\\"3939573a-12a0-436f-b0e5-ab2872a1fde9\\\",\\\"pipe_id\\\":\\\"9db9d8a6-d9e6-4b32-990e-40f0afe4655d\\\",\\\"query\\\":\\\"* | select ack_pps, device_type as type\\\",\\\"sort\\\":\\\"desc\\\",\\\"from\\\":${date_from},\\\"to\\\":${date_to}}     ```      转义前：     ```json     {         \"dataspace_id\":\"3939573a-12a0-436f-b0e5-ab2872a1fde9\",         \"pipe_id\":\"9db9d8a6-d9e6-4b32-990e-40f0afe4655d\",         \"query\":\"* | select ack_pps, device_type as type\",         \"sort\":\"desc\",         \"from\": ${date_from},         \"to\": ${date_to}     }     ``` - query_type为`API`时，function传入api method、url、path_params、headers、response_parser（解析API返回值所需，定义label和json_path将返回值解析为二维表格，label为表头，json_path为字段提取路径）   样例：     ```     {\\\"method\\\":\\\"POST\\\",\\\"uri\\\":\\\"/v1/${project_id}/只填写uri/不带域名/xxx\\\",\\\"headers\\\":{\\\"X-Auth-Token\\\":\\\"${project_token}\\\"},\\\"response_parser\\\":{\\\"labels\\\":[\\\"攻击类型\\\",\\\"攻击源\\\",\\\"时间\\\"],\\\"json_path\\\":[\\\"$.data[:].type\\\",\\\"$.data[:].source\\\",\\\"$.data[:].time\\\"]}}     ```      转义前：     ```json     {         \"method\":\"POST\",         \"uri\":\"/v1/${project_id}/只填写uri/不带域名/xxx\",         \"headers\":{             \"X-Auth-Token\": \"${project_token}\"         },         \"response_parser\":{             \"labels\":[                 \"攻击类型\",                 \"攻击源\",                 \"时间\"             ],             \"json_path\":[                 \"$.data[:].type\",                 \"$.data[:].source\",                 \"$.data[:].time\"             ]         }     }     ``` - query_type为`DSL`时，指定index, dsl(转义成字符串), response_parser   样例：     ```     {\\\"index\\\":\\\"index_xxx_*\\\",\\\"dsl\\\":\\\"{\\\\\\\"query\\\\\\\":{\\\\\\\"match_all\\\\\\\":{}}}\\\",\\\"response_parser\\\":{\\\"labels\\\":[\\\"攻击类型\\\",\\\"攻击源\\\",\\\"时间\\\"],\\\"json_path\\\":[\\\"$.data[:].type\\\",\\\"$.data[:].source\\\",\\\"$.data[:].time\\\"]}}     ```      转义前：      ```json     {         \"index\":\"index_xxx_*\",         \"dsl\":\"{\\\"query\\\":{\\\"match_all\\\":{}}}\",         \"response_parser\":{             \"labels\":[                 \"攻击类型\",                 \"攻击源\",                 \"时间\"             ],             \"json_path\":[                 \"$.data[:].type\",                 \"$.data[:].source\",                 \"$.data[:].time\"             ]         }     }     ``` - query_type为`sql`时，指定opendistro sql插件查询json(转义成字符串)   样例：     ```    {\\\"query\\\":\\\"SELECT count(1) as count , msg.DstPort FROM isap_log_nip_ttl* where oct >= TIMESTAMP(\\\\\\\"${date_from}\\\\\\\") and oct <= TIMESTAMP(\\\\\\\"${date_to}\\\\\\\") group by msg.DstPort order by count desc limit 5\\\"}     ```           转义前：      ```json     {         \"query\":\"SELECT count(1) as count , msg.DstPort FROM isap_log_nip_ttl* where oct >= TIMESTAMP(\\\"${date_from}\\\") and oct <= TIMESTAMP(\\\"${date_to}\\\") group by msg.DstPort order by count desc limit 5\"     }     ```
	QueryFunction string `json:"query_function"`
}

func (o MetricMetaDataDerivedMetrics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricMetaDataDerivedMetrics struct{}"
	}

	return strings.Join([]string{"MetricMetaDataDerivedMetrics", string(data)}, " ")
}

type MetricMetaDataDerivedMetricsQueryType struct {
	value string
}

type MetricMetaDataDerivedMetricsQueryTypeEnum struct {
	CBSL MetricMetaDataDerivedMetricsQueryType
	API  MetricMetaDataDerivedMetricsQueryType
	DSL  MetricMetaDataDerivedMetricsQueryType
	SQL  MetricMetaDataDerivedMetricsQueryType
}

func GetMetricMetaDataDerivedMetricsQueryTypeEnum() MetricMetaDataDerivedMetricsQueryTypeEnum {
	return MetricMetaDataDerivedMetricsQueryTypeEnum{
		CBSL: MetricMetaDataDerivedMetricsQueryType{
			value: "cbsl",
		},
		API: MetricMetaDataDerivedMetricsQueryType{
			value: " api",
		},
		DSL: MetricMetaDataDerivedMetricsQueryType{
			value: " dsl",
		},
		SQL: MetricMetaDataDerivedMetricsQueryType{
			value: " sql",
		},
	}
}

func (c MetricMetaDataDerivedMetricsQueryType) Value() string {
	return c.value
}

func (c MetricMetaDataDerivedMetricsQueryType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MetricMetaDataDerivedMetricsQueryType) UnmarshalJSON(b []byte) error {
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
