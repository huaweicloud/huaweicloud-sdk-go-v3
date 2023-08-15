package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateEdgeAppVersionDto struct {

	// **参数说明**：应用描述。  **取值范围**：只允许中文、字母、数字、下划线（_）、中文分号（；）、中文冒号（：）、中文问号（？）、中文感叹号（！）中文逗号（，）、中文句号（。）、英文引号（;）、英文冒号（:）、英文逗号（,）、英文句号（.）、英文问号（?）、英文感叹号（!）、顿号（、）、连接符（-）的组合。
	Description *string `json:"description,omitempty"`

	ContainerSettings *ContainerSettingsDto `json:"container_settings"`

	// **参数说明**：启动命令。
	Command *interface{} `json:"command,omitempty"`

	// **参数说明**：启动参数。
	Args *interface{} `json:"args,omitempty"`
}

func (o UpdateEdgeAppVersionDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEdgeAppVersionDto struct{}"
	}

	return strings.Join([]string{"UpdateEdgeAppVersionDto", string(data)}, " ")
}
