package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateEsListenerRequestBody struct {
	Listener *EsListenerRequest `json:"listener"`

	// 类型：searchTool 表示修改Elasticsearch/Opensearch负载均衡器，viewTool 表示修改Kibana/Opensearch Dashboard 负载均衡器，默认为searchTool 。
	Type *string `json:"type,omitempty"`
}

func (o UpdateEsListenerRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEsListenerRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateEsListenerRequestBody", string(data)}, " ")
}
