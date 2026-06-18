package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRepositoryNavigationReferencesResponse Response Object
type ListRepositoryNavigationReferencesResponse struct {

	// **参数解释：** 结果标识。 **约束限制：** 不涉及。
	Result *string `json:"result,omitempty"`

	// **参数解释：** 结果消息。 **约束限制：** 不涉及。
	Message *string `json:"message,omitempty"`

	// **参数解释：** def信息。 **约束限制：** 不涉及。
	Defs *[]DefEntryDto `json:"defs,omitempty"`

	// **参数解释：** 索引信息列表。 **约束限制：** 不涉及。
	Refs           *[]RefEntryDto `json:"refs,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListRepositoryNavigationReferencesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRepositoryNavigationReferencesResponse struct{}"
	}

	return strings.Join([]string{"ListRepositoryNavigationReferencesResponse", string(data)}, " ")
}
