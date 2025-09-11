package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NamespaceAllowedEmptyResp **参数解释**： 查询服务的命名空间，各服务命名空间请参考“[服务命名名称](ces_03_0059.xml)”。    **取值范围**： 格式为service.item；service和item必须是字符串，必须以字母开头，只能包含0-9/a-z/A-Z/_。字符串的长度为[0,32]个字符。
type NamespaceAllowedEmptyResp struct {
}

func (o NamespaceAllowedEmptyResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NamespaceAllowedEmptyResp struct{}"
	}

	return strings.Join([]string{"NamespaceAllowedEmptyResp", string(data)}, " ")
}
