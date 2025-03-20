package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Path 资源路径，默认为空串。由若干段字符串拼接而成，每段先包含一个或多个字母、数字、\".\"、\",\"、\"+\"、\"@\"、\"=\"、\"_\"或\"-\"，并以\"/\"结尾，例如\"foo/bar/\"。
type Path struct {
}

func (o Path) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Path struct{}"
	}

	return strings.Join([]string{"Path", string(data)}, " ")
}
