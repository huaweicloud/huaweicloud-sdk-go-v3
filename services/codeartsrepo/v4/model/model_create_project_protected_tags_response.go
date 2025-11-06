package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateProjectProtectedTagsResponse Response Object
type CreateProjectProtectedTagsResponse struct {

	// **参数解释：** 保护tag名称。 **取值范围** 字符串长度不少于1，不超过1000。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 事件列表。
	Actions        *[]ProjectProtectedTagActionDto `json:"actions,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o CreateProjectProtectedTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProjectProtectedTagsResponse struct{}"
	}

	return strings.Join([]string{"CreateProjectProtectedTagsResponse", string(data)}, " ")
}
