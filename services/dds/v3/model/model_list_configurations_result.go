package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListConfigurationsResult struct {

	// 参数模板ID。
	Id string `json:"id"`

	// 参数模板名称。
	Name string `json:"name"`

	// 参数模板描述。
	Description string `json:"description"`

	// 数据库版本。
	DatastoreVersion string `json:"datastore_version"`

	// 数据库类型。
	DatastoreName string `json:"datastore_name"`

	// 参数模板节点类型。 - mongos，表示集群mongos节点类型。 - shard，表示集群shard节点类型。 - config，表示集群config节点类型。 - replica，表示副本集类型。 - single，表示单节点类型。
	NodeType string `json:"node_type"`

	// 创建时间，格式为\"yyyy-MM-ddTHH:mm:ssZ\"。其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	Created string `json:"created"`

	// 更新时间，格式为\"yyyy-MM-ddTHH:mm:ssZ\"。其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	Updated string `json:"updated"`

	// 是否是用户自定义参数模板。 - false表示为默认参数模板。 - true表示为用户自定义参数模板。
	UserDefined bool `json:"user_defined"`
}

func (o ListConfigurationsResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigurationsResult struct{}"
	}

	return strings.Join([]string{"ListConfigurationsResult", string(data)}, " ")
}
