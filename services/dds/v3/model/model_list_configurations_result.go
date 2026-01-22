package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListConfigurationsResult struct {

	// **参数解释：** 参数模板ID。 **取值范围：** 不涉及。
	Id string `json:"id"`

	// **参数解释：** 参数模板名称。 **取值范围：** 不涉及。
	Name string `json:"name"`

	// **参数解释：** 参数模板描述。 **取值范围：** 不涉及。
	Description string `json:"description"`

	// **参数解释：** 数据库版本。 **取值范围：** 不涉及。
	DatastoreVersion string `json:"datastore_version"`

	// **参数解释：** 数据库类型。 **取值范围：** 不涉及。
	DatastoreName string `json:"datastore_name"`

	// **参数解释：** 参数模板节点类型。 **取值范围：** - mongos，表示集群mongos节点类型。 - shard，表示集群shard节点类型。 - config，表示集群config节点类型。 - replica，表示副本集类型。 - readonly，表示副本集只读节点类型。 - shard_readonly，表示集群只读节点类型。 - single，表示单节点类型。
	NodeType string `json:"node_type"`

	// **参数解释：** 创建时间，格式为“yyyy-MM-ddTHH:mm:ssZ”。 其中，T指某个时间的开始；Z指时区偏移量。 **取值范围：** 不涉及。
	Created string `json:"created"`

	// **参数解释：** 更新时间，格式为“yyyy-MM-ddTHH:mm:ssZ”。 其中，T指某个时间的开始；Z指时区偏移量。 **取值范围：** 不涉及。
	Updated string `json:"updated"`

	// **参数解释：** 是否是用户自定义参数模板。 **取值范围：** - 取值为“false”，表示为系统默认参数模板。 - 取值为“true”，表示为用户自定义参数模板。
	UserDefined bool `json:"user_defined"`
}

func (o ListConfigurationsResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigurationsResult struct{}"
	}

	return strings.Join([]string{"ListConfigurationsResult", string(data)}, " ")
}
