package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EngineGroupsInfo struct {
	// 引擎id。

	Id *string `json:"id,omitempty"`
	// 引擎名称。

	Name *string `json:"name,omitempty"`
	// 引擎版本。

	Version *string `json:"version,omitempty"`
	// 可用区列表。

	SupportAzs *[]SupportAzsInfo `json:"supportAzs,omitempty"`
}

func (o EngineGroupsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EngineGroupsInfo struct{}"
	}

	return strings.Join([]string{"EngineGroupsInfo", string(data)}, " ")
}
