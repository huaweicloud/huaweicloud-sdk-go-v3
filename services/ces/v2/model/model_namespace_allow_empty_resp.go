package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NamespaceAllowEmptyResp **参数解释**： 各服务命名空间，请参考“[服务命名空间](ces_03_0059.xml)”。 **取值范围**： 格式为service.item；service和item必须是字符串，必须以字母开头，只能包含0-9/a-z/A-Z/_。长度为[0,32]个字符。
type NamespaceAllowEmptyResp struct {
}

func (o NamespaceAllowEmptyResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NamespaceAllowEmptyResp struct{}"
	}

	return strings.Join([]string{"NamespaceAllowEmptyResp", string(data)}, " ")
}
