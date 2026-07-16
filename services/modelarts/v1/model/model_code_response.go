package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CodeResponse **参数解释：** 代码挂载配置。
type CodeResponse struct {

	// **参数解释：** 代码来源类别。 **取值范围：** - OBS：对象存储服务。 - OBSFS：OBS的文件系统接口。 - EFS：弹性文件服务。
	Source string `json:"source"`

	// **参数解释：** 代码来源地址，格式遵循不同存储系统。 **取值范围：** 不涉及。
	Address *string `json:"address,omitempty"`

	// **参数解释：** 代码来源ID，与address二选一，当且仅当source为EFS时，可以传入sfs turbo的ID。 **取值范围：** 不涉及。
	SourceId *string `json:"source_id,omitempty"`

	// **参数解释：** 挂载到容器内的路径。 **约束限制：** 不涉及。 **取值范围：** 以(/)开头和结尾，可包含字母、数字、中划线、下划线，整个挂载路径长度不能超过255位。 **默认取值：** 不涉及。
	MountPath string `json:"mount_path"`

	// **参数解释：** EFS子路径。 **取值范围：** 不涉及。
	EfsSubPath *string `json:"efs_sub_path,omitempty"`

	// **参数解释：** 挂载权限设置, 是否只读。 **取值范围：** - true：只读。 - false：非只读。
	ReadOnly *bool `json:"read_only,omitempty"`
}

func (o CodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CodeResponse struct{}"
	}

	return strings.Join([]string{"CodeResponse", string(data)}, " ")
}
