package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NamespaceResp **参数解释**： 服务指标命名空间。如：弹性云服务器的命名空间为SYS.ECS，文档数据库的命名空间为SYS.DDS，各服务的命名空间可查看：“[服务命名空间](ces_03_0059.xml)”。 **取值范围**： 格式为service.item；service和item必须是字符串，必须以字母开头，只能包含0-9/a-z/A-Z/_，字符总长度最短为3，最大为32。
type NamespaceResp struct {
}

func (o NamespaceResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NamespaceResp struct{}"
	}

	return strings.Join([]string{"NamespaceResp", string(data)}, " ")
}
