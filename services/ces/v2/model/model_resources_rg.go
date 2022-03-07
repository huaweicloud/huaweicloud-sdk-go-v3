package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourcesRg struct {
	Namespace *Namespace `json:"namespace,omitempty"`
	// 资源实例的维度信息

	Dimensions *[]Dimension `json:"dimensions,omitempty"`
}

func (o ResourcesRg) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourcesRg struct{}"
	}

	return strings.Join([]string{"ResourcesRg", string(data)}, " ")
}
