package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ObjectReference struct {

	// **参数解释**： 资源对象的API类型，例如，DaemonSet、Deployment 等。 **取值范围**： 不涉及。
	Kind *string `json:"kind,omitempty"`

	// **参数解释**： 资源对象的API版本。 **取值范围**： 不涉及。
	ApiVersion *string `json:"apiVersion,omitempty"`

	// **参数解释**： 资源对象的命名空间。 **取值范围**： 不涉及。
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释**： 资源对象的名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 资源对象的唯一标识符（UID）。 **取值范围**： 不涉及。
	Uid *string `json:"uid,omitempty"`

	// **参数解释**： 资源对象的当前版本。 **取值范围**： 不涉及。
	ResourceVersion *string `json:"resourceVersion,omitempty"`
}

func (o ObjectReference) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObjectReference struct{}"
	}

	return strings.Join([]string{"ObjectReference", string(data)}, " ")
}
