package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterFlavorResp **参数解释**： 规格信息。 **取值范围**： 不涉及。
type ClusterFlavorResp struct {

	// **参数解释**： 规格ID。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 链接信息。 **取值范围**： 不涉及。
	Links *[]LinkResp `json:"links,omitempty"`
}

func (o ClusterFlavorResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterFlavorResp struct{}"
	}

	return strings.Join([]string{"ClusterFlavorResp", string(data)}, " ")
}
