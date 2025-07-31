package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NamespaceAllowEmpty 查询服务的命名空间，各服务命名空间请参考“[服务命名空间](ces_03_0059.xml)”
type NamespaceAllowEmpty struct {
}

func (o NamespaceAllowEmpty) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NamespaceAllowEmpty struct{}"
	}

	return strings.Join([]string{"NamespaceAllowEmpty", string(data)}, " ")
}
