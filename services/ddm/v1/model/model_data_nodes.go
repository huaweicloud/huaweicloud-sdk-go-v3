package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DataNodes struct {

	// 后端DN的id。
	Id *string `json:"id,omitempty"`

	// 后端DN的状态。
	Status *string `json:"status,omitempty"`

	// 后端DN的名称。
	Name *string `json:"name,omitempty"`

	// 后端DN的引擎名称。
	EngineName *string `json:"engine_name,omitempty"`

	// 后端DN的引擎版本。
	EngineVersion *string `json:"engine_version,omitempty"`

	// 后端DN的内存大小。
	Mem *int32 `json:"mem,omitempty"`

	// 后端DN的CPU大小。
	Cpu *int32 `json:"cpu,omitempty"`
}

func (o DataNodes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataNodes struct{}"
	}

	return strings.Join([]string{"DataNodes", string(data)}, " ")
}
