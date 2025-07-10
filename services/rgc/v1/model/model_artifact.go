package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Artifact struct {
	Content *Content `json:"content,omitempty"`

	// 策略类型。
	Type *string `json:"type,omitempty"`
}

func (o Artifact) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Artifact struct{}"
	}

	return strings.Join([]string{"Artifact", string(data)}, " ")
}
