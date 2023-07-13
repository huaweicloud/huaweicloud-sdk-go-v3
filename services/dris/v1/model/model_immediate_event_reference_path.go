package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImmediateEventReferencePath struct {

	// **参数说明**：通过点集合定义一个有向的作用范围。
	ActivePath []ImmediateEventPosition3D `json:"active_path"`

	// **参数说明**：事件的影响区域半径，可选，单位为分米。用半径表示影响区域边界离中心线的垂直距离，反映该区域的宽度以覆盖实际路段。
	PathRadius *int32 `json:"path_radius,omitempty"`
}

func (o ImmediateEventReferencePath) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImmediateEventReferencePath struct{}"
	}

	return strings.Join([]string{"ImmediateEventReferencePath", string(data)}, " ")
}
