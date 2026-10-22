package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// V2TagResource **参数解释**：标签资源响应体。
type V2TagResource struct {

	// **参数解释**：标签资源ID。 **取值范围**：由英文字母、数字和中划线组成，长度不超过64个字符。
	ResourceId *string `json:"resource_id,omitempty"`

	// **参数解释**：标签资源详情。 资源对象，用于扩展。默认为空。
	ResourceDetail *interface{} `json:"resource_detail,omitempty"`

	// **参数解释**：标签资源名称。 **取值范围**：不涉及。
	ResourceName *string `json:"resource_name,omitempty"`

	// **参数解释**：标签列表，没有标签默认为空数组，每个元素为一个ResourceTag对象，包含key（标签键）和value（标签值）字段。
	Tags *[]ResourceTag `json:"tags,omitempty"`

	// **参数解释**：系统标签列表，没有系统标签默认不返回，每个元素为一个SystemTag对象，包含key（系统标签键）和value（系统标签值）字段。
	SysTags *[]SystemTag `json:"sys_tags,omitempty"`
}

func (o V2TagResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "V2TagResource struct{}"
	}

	return strings.Join([]string{"V2TagResource", string(data)}, " ")
}
