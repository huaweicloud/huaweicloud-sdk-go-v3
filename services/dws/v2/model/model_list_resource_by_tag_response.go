package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceByTagResponse Response Object
type ListResourceByTagResponse struct {

	// **参数解释**： 标签。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Count *int32 `json:"count,omitempty"`

	// **参数解释**： 资源信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Resources      *[]TagFilter `json:"resources,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListResourceByTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceByTagResponse struct{}"
	}

	return strings.Join([]string{"ListResourceByTagResponse", string(data)}, " ")
}
