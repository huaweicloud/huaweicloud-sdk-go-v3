package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Build 构建信息。
type Build struct {
	Archive *Archive `json:"archive"`

	// 构建附加参数。 - base_image：基础镜像地址。 - build_cmd：自定义构建命令。 - dockerfile_path：自定义dockerfile文件路径 - dockerfile_content：自定义dockerfile内容 - artifact_name: 针对java多模块构建，指定构建后运行的产物，以\".jar\"结尾。
	Parameters map[string]string `json:"parameters"`
}

func (o Build) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Build struct{}"
	}

	return strings.Join([]string{"Build", string(data)}, " ")
}
