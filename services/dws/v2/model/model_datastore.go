package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Datastore **参数解释**： 集群版本。 **取值范围**： 不涉及。
type Datastore struct {

	// **参数解释**： 集群类型。 **取值范围**： 不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**： 集群版本。 **取值范围**： 不涉及。
	Version *string `json:"version,omitempty"`
}

func (o Datastore) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Datastore struct{}"
	}

	return strings.Join([]string{"Datastore", string(data)}, " ")
}
