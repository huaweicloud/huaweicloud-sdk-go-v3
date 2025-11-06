package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDataStoreFlavorDetailRequest Request Object
type ShowDataStoreFlavorDetailRequest struct {

	// **参数解释**： 引擎类型id。 **约束限制**： 不涉及 **取值范围**： Elasticsearch：cf7e2c8f-096c-4fcf-b174-1ebe060679fb。 Opensearch：07ec9f86-ec2f-49e7-8913-373003aedf32。 Logstash: 575276bb-87e5-4e18-8e1e-e748d8ad3a06。 **默认取值**： 不涉及
	DatastoreId string `json:"datastore_id"`

	// **参数解释**： 引擎类型id。 **约束限制**： 不涉及 **取值范围**： Elasticsearch 7.10.2：01f53413-0a58-4b0c-848a-f625846bae23。 Opensearch 2.19.0：11a9df5c-711f-496c-866d-a4521c179671。 Logstash 7.10.0: f5609cf0-3514-49ef-87db-a3df2858a46f。 **默认取值**： 不涉及
	DatastoreVersionId *string `json:"datastore_version_id,omitempty"`
}

func (o ShowDataStoreFlavorDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDataStoreFlavorDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowDataStoreFlavorDetailRequest", string(data)}, " ")
}
