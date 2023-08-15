package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EventReferencePath struct {

	// **参数说明**：激活路径。
	ActivePath *[]EventLocation `json:"active_path,omitempty"`

	// **参数说明**：事件的影响区域半径，可选，单位为分米。用半径表示影响区域边界离中心线的垂直距离，反映该区域的宽度以覆盖实际路段。
	PathRadius *int32 `json:"path_radius,omitempty"`
}

func (o EventReferencePath) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventReferencePath struct{}"
	}

	return strings.Join([]string{"EventReferencePath", string(data)}, " ")
}
