package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FileInfo **参数解释：** 模型&权重&代码文件存储挂载配置。 **约束限制：** 不涉及。
type FileInfo struct {

	// **参数解释：** 代码来源类别。 **约束限制：** 不涉及。 **取值范围：** 如下参数不区分大小写 - OBS：对象存储服务。 - OBSFS：OBS的文件系统接口。 - EFS：弹性文件服务。 - LOCAL：挂载宿主机本地存储目录。 **默认取值：** 不涉及。
	Source string `json:"source"`

	// **参数解释：** 代码来源地址，格式遵循不同存储系统。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Address *string `json:"address,omitempty"`

	// **参数解释：** 挂载到容器内的路径。 **约束限制：** 不涉及。 **取值范围：** 以(/)开头和结尾，可包含字母、数字、中划线、下划线，整个挂载路径长度不能超过255位。 **默认取值：** 不涉及。
	MountPath string `json:"mount_path"`

	// **参数解释：** 是否支持模型本地缓存，默认是不支持。 **约束限制：** 不涉及。 **取值范围：** - true：支持。 - false：不支持。 **默认取值：** false。
	HostCache *bool `json:"host_cache,omitempty"`

	// **参数解释：** 当存储类别为EFS时，支持配置子目录。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	EfsSubPath *string `json:"efs_sub_path,omitempty"`

	// **参数解释：** 挂载权限设置，是否只读。 **约束限制：** 不涉及。 **取值范围：** - true：只读。 - false：非只读。 **默认取值：** 不涉及。
	ReadOnly *bool `json:"read_only,omitempty"`

	// **参数解释：** OS预热。 **约束限制：** 不涉及。 **取值范围：** - true：预热。 - false：不预热。 **默认取值：** 不涉及。
	OsWarmUp *bool `json:"os_warm_up,omitempty"`

	// **参数解释：** 预热名称。 **约束限制：** os_warm_up为true时必填。 **取值范围：** 支持1-64位字符，可包含字母、中文、数字、中划线、下划线。 **默认取值：** 不涉及。
	SourceName *string `json:"source_name,omitempty"`
}

func (o FileInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FileInfo struct{}"
	}

	return strings.Join([]string{"FileInfo", string(data)}, " ")
}
