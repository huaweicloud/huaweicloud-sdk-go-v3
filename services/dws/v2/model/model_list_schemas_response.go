package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSchemasResponse Response Object
type ListSchemasResponse struct {

	// **参数解释**： 集群模式空间信息列表。 **取值范围**： 不涉及。
	Schemas *[]SchemaInfo `json:"schemas,omitempty"`

	// **参数解释**： 总数量。 **取值范围**： 不涉及。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSchemasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSchemasResponse struct{}"
	}

	return strings.Join([]string{"ListSchemasResponse", string(data)}, " ")
}
