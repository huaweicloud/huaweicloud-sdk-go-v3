package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ActionOnComponentBuild 构建信息。
type ActionOnComponentBuild struct {
	Archive *Archive `json:"archive,omitempty"`

	// 构建附加参数。 - base_image：基础镜像地址。 - build_cmd：自定义构建命令。 - dockerfile_path：自定义dockerfile文件路径 - dockerfile_content：自定义dockerfile内容
	Parameters map[string]string `json:"parameters,omitempty"`
}

func (o ActionOnComponentBuild) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionOnComponentBuild struct{}"
	}

	return strings.Join([]string{"ActionOnComponentBuild", string(data)}, " ")
}
