package model

import (
	"encoding/json"

	"strings"
)

type ComponentList struct {
	//   组件名称  MRS 2.1.0及之后版本支持Presto、Hadoop、Spark、HBase、Hive、Tez、Hue、Loader、Flink、Impala、Kudu、Flume、Kafka和Storm组件。 MRS 2.0.5版本支持Presto、Hadoop、Spark、HBase、Hive、Tez、Hue、Loader、Flume、Kafka和Storm组件。 MRS 1.8.10版本支持Presto、Hadoop、Spark、HBase、Opentsdb、Hive、Hue、Loader、Flink、Flume、Kafka、KafkaManager和Storm组件。

	ComponentName string `json:"component_name"`
}

func (o ComponentList) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ComponentList struct{}"
	}

	return strings.Join([]string{"ComponentList", string(data)}, " ")
}
