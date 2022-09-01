package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 卷配置
type Volumes struct {

	// 卷名称，小写字母或数字，最长63个字符
	Name string `json:"name" xml:"name"`

	// 卷的类型，支持configMap,secret,emptyDir,hostPath
	Type string `json:"type" xml:"type"`

	// 卷来源，type为hostPath时输入路径，要求以/开头，后面可包含中划线，反斜杠，下划线，点号，字母，数字； secret时输入secret名称，configMap时输入configMap名称，emptyDir时输入disk或memory
	Source string `json:"source" xml:"source"`

	// 卷挂载路径，必须是合法的路径
	Destination string `json:"destination" xml:"destination"`

	// 读写权限，configMap和secret类型只支持读权限
	ReadOnly bool `json:"read_only" xml:"read_only"`
}

func (o Volumes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Volumes struct{}"
	}

	return strings.Join([]string{"Volumes", string(data)}, " ")
}
