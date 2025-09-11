package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NamespaceResp **参数解释**： 查询服务的命名空间，各服务命名空间请参考“[服务命名空间](ces_03_0059.xml)”。 **取值范围**： 格式为service.item；service和item必须是字符串，必须以字母开头，只能包含0-9/a-z/A-Z/_。字符串的长度必须在 3 到 32个字符之间。
type NamespaceResp struct {
}

func (o NamespaceResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NamespaceResp struct{}"
	}

	return strings.Join([]string{"NamespaceResp", string(data)}, " ")
}
