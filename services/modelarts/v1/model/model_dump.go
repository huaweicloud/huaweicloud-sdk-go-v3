package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Dump **参数解释：** 用户产物转储配置 **约束限制：** 不涉及。
type Dump struct {

	// **参数解释：** 转储挂载目录来源类别。 **约束限制：** 不涉及。 **取值范围：** - [OBS：对象存储服务。](tag:hws,hws_hk) - OBSFS：OBS的文件系统接口。 **默认取值：** 不涉及。
	Source string `json:"source"`

	// **参数解释：** 转储挂载目录来源地址，支持https或obs协议的obs地址。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Address *string `json:"address,omitempty"`

	// **参数解释：** 挂载到容器内的路径，要求以/开头，后面可包含中划线，反斜杠，下划线，点号，字母，数字。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	MountPath string `json:"mount_path"`
}

func (o Dump) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Dump struct{}"
	}

	return strings.Join([]string{"Dump", string(data)}, " ")
}
