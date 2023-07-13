package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MotionItem struct {

	// 时间戳，相对时间戳，单位S，保留3位小数。
	Timestamp *float32 `json:"timestamp,omitempty"`

	// root 3维坐标。
	Root *[]interface{} `json:"root,omitempty"`

	// 75个关节点,四元数。
	Joints *[]interface{} `json:"joints,omitempty"`
}

func (o MotionItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MotionItem struct{}"
	}

	return strings.Join([]string{"MotionItem", string(data)}, " ")
}
