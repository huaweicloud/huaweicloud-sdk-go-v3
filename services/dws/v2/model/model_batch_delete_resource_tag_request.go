package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteResourceTagRequest Request Object
type BatchDeleteResourceTagRequest struct {

	// **参数解释**： 集群ID。获取方式方法请参见[获取集群ID](dws_02_00068.xml)。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ClusterId string `json:"cluster_id"`

	Body *BatchDeleteResourceTags `json:"body,omitempty"`
}

func (o BatchDeleteResourceTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteResourceTagRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteResourceTagRequest", string(data)}, " ")
}
