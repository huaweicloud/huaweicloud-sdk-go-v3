package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AccessScope struct {

	// **参数解释：** 集群命名空间的列表，允许使用通配符（“\\*”），表示所有命名空间。当选择了不同集群时，命名空间的列表可以为多个集群的集合，在转化成RBAC策略时，会自动判断集群下的命名空间是否存在。 **约束限制：** 当前最多支持同时授权500个命名空间。 **取值范围：** \\[\"\\*\"\\]或者集群命名空间列表。 **默认取值：** 不涉及
	Namespaces []string `json:"namespaces"`
}

func (o AccessScope) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessScope struct{}"
	}

	return strings.Join([]string{"AccessScope", string(data)}, " ")
}
