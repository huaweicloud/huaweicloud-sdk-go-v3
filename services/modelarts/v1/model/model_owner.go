package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Owner 桶拥有者信息，包含租户ID。
type Owner struct {

	// **参数解释**：对象拥有者的DomainId。 **取值范围**：不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：对象拥有者的名称。 **取值范围**：不涉及。
	DisplayName *string `json:"display_name,omitempty"`
}

func (o Owner) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Owner struct{}"
	}

	return strings.Join([]string{"Owner", string(data)}, " ")
}
