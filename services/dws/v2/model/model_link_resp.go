package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LinkResp **参数解释**： 链接信息。 **取值范围**： 不涉及。
type LinkResp struct {

	// **参数解释**： 关联信息。 **取值范围**： 不涉及。
	Rel *string `json:"rel,omitempty"`

	// **参数解释**： 链接信息。 **取值范围**： 不涉及。
	Href *string `json:"href,omitempty"`
}

func (o LinkResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LinkResp struct{}"
	}

	return strings.Join([]string{"LinkResp", string(data)}, " ")
}
