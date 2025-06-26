package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DatabaseObjectInfo **参数解释**： 对象信息。 **取值范围**： 不涉及。
type DatabaseObjectInfo struct {

	// **参数解释**： 对象名称。 **取值范围**： 不涉及。
	ObjName *string `json:"obj_name,omitempty"`
}

func (o DatabaseObjectInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabaseObjectInfo struct{}"
	}

	return strings.Join([]string{"DatabaseObjectInfo", string(data)}, " ")
}
