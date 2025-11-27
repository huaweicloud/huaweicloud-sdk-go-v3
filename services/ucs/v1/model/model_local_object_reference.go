package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LocalObjectReference struct {

	// 被引用资源的名称
	Name *string `json:"name,omitempty"`
}

func (o LocalObjectReference) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LocalObjectReference struct{}"
	}

	return strings.Join([]string{"LocalObjectReference", string(data)}, " ")
}
