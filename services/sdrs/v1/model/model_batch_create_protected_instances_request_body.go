package model

import (
	"encoding/json"

	"strings"
)

// 批量创建保护实例请求数据接口
type BatchCreateProtectedInstancesRequestBody struct {
	ProtectedInstances *BatchCreateProtectedInstancesRequestParams `json:"protected_instances"`
}

func (o BatchCreateProtectedInstancesRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchCreateProtectedInstancesRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateProtectedInstancesRequestBody", string(data)}, " ")
}
