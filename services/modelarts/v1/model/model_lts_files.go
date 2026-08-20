package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LtsFiles **参数解释：**  容器日志文件配置。 **约束限制：**  数量上限为10个。
type LtsFiles struct {

	// **参数解释：** 日志文件路径。 **约束限制：** 1.路径必须以 / 开头，且第一级目录不能使用通配符，只能包含大写字母，小写字母，数字或特殊符号-_/_*?，长度不能超过 512 个字符。 2.最多允许三级目录使用通配符进行匹配。 **取值范围：** 不涉及 **默认取值：** 不涉及。
	LogPath string `json:"log_path"`

	// **参数解释：** 日志文件名称。 **约束限制：** 只能包含大写字母，小写字母，数字或特殊字符-_*?，不支持.gz .tar .zip后缀类型，长度不能超过 255 个字符。 **取值范围：** 不涉及 **默认取值：** 不涉及。
	FilePattern string `json:"file_pattern"`
}

func (o LtsFiles) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LtsFiles struct{}"
	}

	return strings.Join([]string{"LtsFiles", string(data)}, " ")
}
