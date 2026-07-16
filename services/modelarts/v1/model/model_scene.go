package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Scene 场景。
type Scene struct {

	// 场景ID。
	Id *string `json:"id,omitempty"`

	// 场景名称。
	Name *string `json:"name,omitempty"`

	// 节点列表。
	Steps *[]string `json:"steps,omitempty"`
}

func (o Scene) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Scene struct{}"
	}

	return strings.Join([]string{"Scene", string(data)}, " ")
}
