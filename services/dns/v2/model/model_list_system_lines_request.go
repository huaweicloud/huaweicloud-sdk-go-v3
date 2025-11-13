package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSystemLinesRequest Request Object
type ListSystemLinesRequest struct {

	// **参数解释：** 指定显示语言。 **约束限制：** 不涉及。 **取值范围：** - zh-cn：中文 - en-us：英语             - es-us：西班牙语 - pt-br：葡萄牙语 **默认取值：** zh-cn
	Locale *string `json:"locale,omitempty"`

	// **参数解释：** 分页查询时配置每页返回的资源个数。 **约束限制：** 不涉及。 **取值范围：** 0~1000。 **默认取值：** 1000
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 分页查询起始偏移量，表示从偏移量的下一个资源开始查询。 **约束限制：** 当设置marker不为空时，以marker为分页起始标识，offset不生效。 **取值范围：** 0~2147483647。 **默认取值：** 0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListSystemLinesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSystemLinesRequest struct{}"
	}

	return strings.Join([]string{"ListSystemLinesRequest", string(data)}, " ")
}
