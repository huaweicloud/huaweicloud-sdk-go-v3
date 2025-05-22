package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlavorInfoResponse **参数解释**： 规格详情。 **取值范围**： 不涉及。
type FlavorInfoResponse struct {

	// **参数解释**： 规格ID。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 规格编码。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： CPU数。 **取值范围**： 不涉及。
	Vcpus *string `json:"vcpus,omitempty"`

	// **参数解释**： 内存数。 **取值范围**： 不涉及。
	Ram *string `json:"ram,omitempty"`

	// **参数解释**： 是否是当前规格。 **取值范围**： 不涉及。
	IsCurrentFlavor *bool `json:"is_current_flavor,omitempty"`
}

func (o FlavorInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlavorInfoResponse struct{}"
	}

	return strings.Join([]string{"FlavorInfoResponse", string(data)}, " ")
}
