package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LabelInfo 标签信息
type LabelInfo struct {

	// 标签名称
	Key *string `json:"key,omitempty"`

	// 标签值
	Val *string `json:"val,omitempty"`
}

func (o LabelInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LabelInfo struct{}"
	}

	return strings.Join([]string{"LabelInfo", string(data)}, " ")
}
