package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LineStatus struct {
	StartPoint *Point `json:"startPoint,omitempty"`

	EndPoint *Point `json:"endPoint,omitempty"`

	// **参数解释：** 表示是否为关键线路（关键线路未执行无法取消升级流程） **约束限制：** 不涉及 **取值范围：** - true：是关键线路 - false：非关键线路  **默认取值：** 不涉及
	Critical *bool `json:"critical,omitempty"`
}

func (o LineStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LineStatus struct{}"
	}

	return strings.Join([]string{"LineStatus", string(data)}, " ")
}
