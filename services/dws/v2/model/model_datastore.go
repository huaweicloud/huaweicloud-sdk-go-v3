package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Datastore 集群版本
type Datastore struct {

	// 集群类型。
	Type *string `json:"type,omitempty"`

	// 集群版本。
	Version *string `json:"version,omitempty"`
}

func (o Datastore) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Datastore struct{}"
	}

	return strings.Join([]string{"Datastore", string(data)}, " ")
}
