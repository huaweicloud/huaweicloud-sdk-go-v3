package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NamespaceSchema 服务指标命名空间，格式为service.item；service和item必须是字符串，必须以字母开头，只能包含0-9/a-z/A-Z/_，字符总长度最短为3，最大为32。说明： 当alarm_type为（EVENT.SYS| EVENT.CUSTOM）时允许为空；如：弹性云服务器的命名空间为SYS.ECS，文档数据库的命名空间为SYS.DDS，各服务的命名空间可查看：“[服务命名空间](ces_03_0059.xml)”。
type NamespaceSchema struct {
}

func (o NamespaceSchema) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NamespaceSchema struct{}"
	}

	return strings.Join([]string{"NamespaceSchema", string(data)}, " ")
}
