package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 应用扩展功能配置选项
type Annotations struct {

	// 生成ranktablefile。该参数目前只支持赋值\"ascend-1980\"，指昇腾D910。
	RingController *string `json:"ring_controller,omitempty"`

	// 离线自愈功能配置字段，须填写调度的节点组id
	AutonomyEdgeSelector *string `json:"autonomy_edge_selector,omitempty"`
}

func (o Annotations) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Annotations struct{}"
	}

	return strings.Join([]string{"Annotations", string(data)}, " ")
}
