package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DatastoreResult struct {

	// node_type参数模板节点类型。取值范围： - mongos，表示集群mongos节点类型。 - shard，表示集群shard节点类型。 - config，表示集群config节点类型。 - replica，表示副本集类型。 - single，表示单节点类型。
	NodeType string `json:"node_type"`

	// version数据库版本。
	Version string `json:"version"`
}

func (o DatastoreResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatastoreResult struct{}"
	}

	return strings.Join([]string{"DatastoreResult", string(data)}, " ")
}
