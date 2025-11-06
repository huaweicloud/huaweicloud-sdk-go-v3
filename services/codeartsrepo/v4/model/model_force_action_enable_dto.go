package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ForceActionEnableDto struct {

	// **参数解释：** 强制操作名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 强制操作是否开启。
	Enable *bool `json:"enable,omitempty"`
}

func (o ForceActionEnableDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ForceActionEnableDto struct{}"
	}

	return strings.Join([]string{"ForceActionEnableDto", string(data)}, " ")
}
