package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttributeOperationDto An operation that applies to the requested group. This operation might add, replace, or remove an attribute.
type AttributeOperationDto struct {

	// 要更新的属性的字符串表示
	AttributePath string `json:"attribute_path"`

	// 要更新的属性值。当属性为对象类型时，此处填写对象的JSON字符串。为null时表示删除对应属性
	AttributeValue *string `json:"attribute_value,omitempty"`
}

func (o AttributeOperationDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttributeOperationDto struct{}"
	}

	return strings.Join([]string{"AttributeOperationDto", string(data)}, " ")
}
