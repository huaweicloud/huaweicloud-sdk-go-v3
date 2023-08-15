package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReferencePath struct {

	// **参数说明**：事件的有效路径
	ActivePath *[]Position3D `json:"active_path,omitempty"`

	// **参数说明**：事件的影响区域半径，可选，单位为分米。用半径表示影响区域边界离中心线的垂直距离，反映该区域的宽度以覆盖实际路段。
	PathRadius *int32 `json:"path_radius,omitempty"`
}

func (o ReferencePath) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReferencePath struct{}"
	}

	return strings.Join([]string{"ReferencePath", string(data)}, " ")
}
