package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ActionsList struct {

	// 置信度，取值范围0～1。
	Confidence float64 `json:"confidence"`

	// 动作编号，取值范围：[1,2,3,4]，其中： • 1：左摇头 • 2：右摇头 • 3：点头 • 4：嘴部动作
	Action *int32 `json:"action,omitempty"`
}

func (o ActionsList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionsList struct{}"
	}

	return strings.Join([]string{"ActionsList", string(data)}, " ")
}
