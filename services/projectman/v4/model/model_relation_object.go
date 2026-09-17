package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RelationObject 关系对应的实体，比如工作项
type RelationObject struct {

	// **参数解释**： 实体的种类。 **取值范围**： 不涉及。
	ObjectType *string `json:"object_type,omitempty"`

	// **参数解释**： 类型列表。 **取值范围**： 不涉及。
	Categories *[]string `json:"categories,omitempty"`
}

func (o RelationObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RelationObject struct{}"
	}

	return strings.Join([]string{"RelationObject", string(data)}, " ")
}
