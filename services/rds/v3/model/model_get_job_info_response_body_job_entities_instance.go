package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetJobInfoResponseBodyJobEntitiesInstance 创建实例，单转主备，创建只读实例时返回。
type GetJobInfoResponseBodyJobEntitiesInstance struct {

	// 实例的连接地址。
	Endpoint *string `json:"endpoint,omitempty"`

	// 实例类型。
	Type *string `json:"type,omitempty"`

	Datastore *GetJobInfoResponseBodyJobEntitiesInstanceDatastore `json:"datastore,omitempty"`

	// 主实例ID，仅创建只读实例的时候返回。
	ReplicaOf *string `json:"replica_of,omitempty"`
}

func (o GetJobInfoResponseBodyJobEntitiesInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetJobInfoResponseBodyJobEntitiesInstance struct{}"
	}

	return strings.Join([]string{"GetJobInfoResponseBodyJobEntitiesInstance", string(data)}, " ")
}
