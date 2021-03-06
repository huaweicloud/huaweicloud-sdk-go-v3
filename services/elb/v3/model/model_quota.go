package model

import (
	"encoding/json"

	"strings"
)

// 查询配额限制去请求返回对象
type Quota struct {
	// 证书配额。 -1表示无配额限制。

	Certificate int32 `json:"certificate"`
	// 健康检查配额。 -1表示无配额限制。

	Healthmonitor int32 `json:"healthmonitor"`
	// 转发策略配额。 -1表示无配额限制。

	L7policy int32 `json:"l7policy"`
	// 监听器配额。 -1表示无配额限制。

	Listener int32 `json:"listener"`
	// 负载均衡器配额。 -1表示无配额限制。

	Loadbalancer int32 `json:"loadbalancer"`
	// 后端云服务器配额。 -1表示无配额限制。

	Member int32 `json:"member"`
	// 后端云服务器组配额。 -1表示无配额限制。

	Pool int32 `json:"pool"`
	// 项目ID。

	ProjectId string `json:"project_id"`
}

func (o Quota) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Quota struct{}"
	}

	return strings.Join([]string{"Quota", string(data)}, " ")
}
