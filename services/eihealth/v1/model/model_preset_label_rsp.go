package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PresetLabelRsp 预置标签
type PresetLabelRsp struct {

	// 标签名称
	Name *string `json:"name,omitempty"`
}

func (o PresetLabelRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PresetLabelRsp struct{}"
	}

	return strings.Join([]string{"PresetLabelRsp", string(data)}, " ")
}
