package model

import (
	"encoding/json"

	"strings"
)

// 转发IoTA服务消息结构
type IoTaForwarding struct {
	// **参数说明**：IoTA服务对应的数据源Id。

	DataSourceId string `json:"data_source_id"`
	// **参数说明**：IoTA服务对应的projectId信息。

	ProjectId string `json:"project_id"`
}

func (o IoTaForwarding) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "IoTaForwarding struct{}"
	}

	return strings.Join([]string{"IoTaForwarding", string(data)}, " ")
}
