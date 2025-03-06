package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceEngineDetail struct {

	// 引擎版本号。
	EngineVersion *string `json:"engine_version,omitempty"`

	// 实例详情。
	Instances *[]InstanceDetail `json:"instances,omitempty"`
}

func (o InstanceEngineDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceEngineDetail struct{}"
	}

	return strings.Join([]string{"InstanceEngineDetail", string(data)}, " ")
}
