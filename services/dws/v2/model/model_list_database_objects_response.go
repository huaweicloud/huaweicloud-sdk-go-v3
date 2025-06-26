package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatabaseObjectsResponse Response Object
type ListDatabaseObjectsResponse struct {

	// **参数解释**： 类型。 **取值范围**： DATABASE、SCHEMA、TABLE、VIEW、COLUMN、FUNCTION、SEQUENCE、NODEGROUP
	Type *string `json:"type,omitempty"`

	// **参数解释**： 对象列表。 **取值范围**： 不涉及。
	ObjectList *[]DatabaseObjectInfo `json:"object_list,omitempty"`

	// **参数解释**： 对象总条数。 **取值范围**： 不涉及。
	Count          *string `json:"count,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListDatabaseObjectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabaseObjectsResponse struct{}"
	}

	return strings.Join([]string{"ListDatabaseObjectsResponse", string(data)}, " ")
}
