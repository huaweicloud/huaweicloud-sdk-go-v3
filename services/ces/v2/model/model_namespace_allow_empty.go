package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NamespaceAllowEmpty **参数解释**： 查询服务的命名空间，各服务命名空间请参考“[服务命名空间](ces_03_0059.xml)”。 **约束限制**： 不涉及。 **取值范围**： 格式为service.item；service和item必须是字符串，必须以字母开头，只能包含0-9/a-z/A-Z/_。字符串的长度必须在 0 到 32个字符之间。 **默认取值**： 不涉及。
type NamespaceAllowEmpty struct {
}

func (o NamespaceAllowEmpty) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NamespaceAllowEmpty struct{}"
	}

	return strings.Join([]string{"NamespaceAllowEmpty", string(data)}, " ")
}
