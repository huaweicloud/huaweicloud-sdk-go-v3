package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EngineGroupsInfo struct {

	// 引擎id。
	Id *string `json:"id,omitempty" xml:"id"`

	// 引擎名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 引擎版本。
	Version *string `json:"version,omitempty" xml:"version"`

	// 可用区列表。
	SupportAzs *[]SupportAzsInfo `json:"supportAzs,omitempty" xml:"supportAzs"`
}

func (o EngineGroupsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EngineGroupsInfo struct{}"
	}

	return strings.Join([]string{"EngineGroupsInfo", string(data)}, " ")
}
