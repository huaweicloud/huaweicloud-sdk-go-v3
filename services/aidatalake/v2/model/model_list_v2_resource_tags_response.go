package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListV2ResourceTagsResponse Response Object
type ListV2ResourceTagsResponse struct {

	// **参数解释**：资源标签列表，每个元素为一个ResourceTag对象，包含key（标签键）和value（标签值）字段。
	Tags *[]ResourceTag `json:"tags,omitempty"`

	// **参数解释**：系统标签列表，每个元素为一个SystemTag对象，包含key（系统标签键）和value（系统标签值）字段。
	SysTags *[]SystemTag `json:"sys_tags,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListV2ResourceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListV2ResourceTagsResponse struct{}"
	}

	return strings.Join([]string{"ListV2ResourceTagsResponse", string(data)}, " ")
}
