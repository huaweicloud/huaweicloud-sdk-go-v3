package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateIgnorePathRequestBody struct {

	// **参数解释：** 屏蔽目录的节点信息列表。 **约束限制：** 不涉及。
	IgnorePathSettings *[]IgnorePathSettingItem `json:"ignore_path_settings,omitempty"`
}

func (o UpdateIgnorePathRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIgnorePathRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateIgnorePathRequestBody", string(data)}, " ")
}
