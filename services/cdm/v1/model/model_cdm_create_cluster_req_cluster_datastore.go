package model

import (
	"encoding/json"

	"strings"
)

// 集群信息，请参见•datastore参数说明
type CdmCreateClusterReqClusterDatastore struct {
	// 集群类型，当前只有“cdm”一种类型

	Type string `json:"type"`
	// 集群版本

	Version string `json:"version"`
}

func (o CdmCreateClusterReqClusterDatastore) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CdmCreateClusterReqClusterDatastore struct{}"
	}

	return strings.Join([]string{"CdmCreateClusterReqClusterDatastore", string(data)}, " ")
}
