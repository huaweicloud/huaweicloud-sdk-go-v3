package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceTagsResponse Response Object
type ListInstanceTagsResponse struct {

	// **参数解释**：  资源标签列表，包含该实例下所有标签的键值对信息。详情请参见BatchResourceTag。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Tags           *[]BatchResourceTag `json:"tags,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListInstanceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceTagsResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceTagsResponse", string(data)}, " ")
}
