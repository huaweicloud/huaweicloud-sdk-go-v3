package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeployKeyParamsDto struct {

	// **参数解释：** 标题。 **取值范围：** 字符串长度不少于1，不超过1000。
	Title *string `json:"title,omitempty"`

	// **参数解释：** key值。 **取值范围：** 字符串长度不少于1，不超过1000。
	Key *string `json:"key,omitempty"`
}

func (o DeployKeyParamsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeployKeyParamsDto struct{}"
	}

	return strings.Join([]string{"DeployKeyParamsDto", string(data)}, " ")
}
