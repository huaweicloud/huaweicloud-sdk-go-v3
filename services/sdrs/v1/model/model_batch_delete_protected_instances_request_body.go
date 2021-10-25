package model

import (
	"encoding/json"

	"strings"
)

// 批量删除请求体
type BatchDeleteProtectedInstancesRequestBody struct {
	// 所需要删除的保护实例列表。

	ProtectedInstances []ResourceId `json:"protected_instances"`
	// 是否删除容灾站点服务器，默认值为false。

	DeleteTargetServer *bool `json:"delete_target_server,omitempty"`
	// 是否删除容灾站点弹性IP，默认值为false。

	DeleteTargetEip *bool `json:"delete_target_eip,omitempty"`
}

func (o BatchDeleteProtectedInstancesRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchDeleteProtectedInstancesRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteProtectedInstancesRequestBody", string(data)}, " ")
}
