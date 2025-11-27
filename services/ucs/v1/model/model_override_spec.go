package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OverrideSpec struct {

	// 资源选择器，限制该覆盖策略适用的资源类型
	ResourceSelectors *[]ResourceSelector `json:"resourceSelectors,omitempty"`

	// 将应用于资源的覆盖规则
	Overriders *interface{} `json:"overriders,omitempty"`
}

func (o OverrideSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OverrideSpec struct{}"
	}

	return strings.Join([]string{"OverrideSpec", string(data)}, " ")
}
