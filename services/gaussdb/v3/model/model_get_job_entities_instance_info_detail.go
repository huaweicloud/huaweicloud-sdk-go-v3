package model

import (
	"encoding/json"

	"strings"
)

// 执行任务的实例信息。
type GetJobEntitiesInstanceInfoDetail struct {
	// 实例的连接地址。

	Endpoint string `json:"endpoint"`
	// 实例类型。

	Type string `json:"type"`

	Datastore *GetJobInstanceDatastoreInfoDetail `json:"datastore"`
}

func (o GetJobEntitiesInstanceInfoDetail) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "GetJobEntitiesInstanceInfoDetail struct{}"
	}

	return strings.Join([]string{"GetJobEntitiesInstanceInfoDetail", string(data)}, " ")
}
