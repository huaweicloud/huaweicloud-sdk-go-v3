package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PluginExecutionVo struct {

	// **参数解释**： 插件执行脚本类型。 **取值范围**： - Shell：Shell类型脚本。 - Nodejs：Nodejs类型脚本。 - Java：Java类型脚本。 - Python：Python类型脚本。
	Type *string `json:"type,omitempty"`

	// **参数解释**： 插件执行脚本入口文件。 **取值范围**： 不涉及。
	Target *string `json:"target,omitempty"`

	// **参数解释**： 插件的OBS存放路径。 **取值范围**： 不涉及。
	FilePath *string `json:"file_path,omitempty"`
}

func (o PluginExecutionVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginExecutionVo struct{}"
	}

	return strings.Join([]string{"PluginExecutionVo", string(data)}, " ")
}
