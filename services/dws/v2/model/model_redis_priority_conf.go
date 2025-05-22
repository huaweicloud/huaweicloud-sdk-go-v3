package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RedisPriorityConf **参数解释**： 待重分布表优先级请求信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type RedisPriorityConf struct {

	// **参数解释**： 数据库名称。 **约束限制**： 不涉及。 **取值范围**： 1~1024 **默认取值**： 不涉及。
	DbName *string `json:"db_name,omitempty"`

	// **参数解释**： 待重分布表优先级请求信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Priority *[]RedisPriorityTable `json:"priority,omitempty"`
}

func (o RedisPriorityConf) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RedisPriorityConf struct{}"
	}

	return strings.Join([]string{"RedisPriorityConf", string(data)}, " ")
}
