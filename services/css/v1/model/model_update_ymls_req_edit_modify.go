package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateYmlsReqEditModify struct {

	// **参数解释**： 参数配置列表，值为需要修改的json数据，OpenSearch集群也使用此参数，即修改opensearch.yml时，这里也是填写elasticsearch.yml。 **约束限制**： 不涉及
	ElasticsearchYml *interface{} `json:"elasticsearch.yml,omitempty"`

	// **参数解释**： 参数配置列表，值为需要修改的json数据。OpenSearch集群也使用此参数，即修改opensearch_dashboards.yml时，这里也是填写kibana.yml。 **约束限制**： 不涉及
	KibanaYml *interface{} `json:"kibana.yml,omitempty"`

	SystemConfig *interface{} `json:"system.config,omitempty"`
}

func (o UpdateYmlsReqEditModify) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateYmlsReqEditModify struct{}"
	}

	return strings.Join([]string{"UpdateYmlsReqEditModify", string(data)}, " ")
}
