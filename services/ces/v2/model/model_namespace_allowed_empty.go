package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NamespaceAllowedEmpty 查询服务的命名空间，各服务命名空间请参考“[服务维度名称](ces_03_0059.xml)”
type NamespaceAllowedEmpty struct {
}

func (o NamespaceAllowedEmpty) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NamespaceAllowedEmpty struct{}"
	}

	return strings.Join([]string{"NamespaceAllowedEmpty", string(data)}, " ")
}
