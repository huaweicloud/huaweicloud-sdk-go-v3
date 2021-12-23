package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// cdm信息
type Datastore struct {
	// 类型，一般为cdm。

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
