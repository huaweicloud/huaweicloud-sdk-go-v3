package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CustomTimeInfo struct {

	// **参数解释：** 是否开启自定义时间字段。 **取值范围：** - true - fasle
	Enable *bool `json:"enable,omitempty"`

	// **参数解释：** 从demoField中选取的作为日志系统时间的字段名称。 **取值范围：** 不涉及。
	Key *string `json:"key,omitempty"`

	// **参数解释：** 从demoField中选取的作为日志系统时间的字段内容示例。 **取值范围：** 不涉及。
	Value *string `json:"value,omitempty"`

	// **参数解释：** 用于解析字段为时间类型的时间格式参数。 **取值范围：** 不涉及。
	TimeFormat *string `json:"time_format,omitempty"`
}

func (o CustomTimeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomTimeInfo struct{}"
	}

	return strings.Join([]string{"CustomTimeInfo", string(data)}, " ")
}
