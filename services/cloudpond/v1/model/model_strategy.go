package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Strategy struct {
	Status *OperationStatus `json:"status,omitempty"`

	Scene *Scene `json:"scene,omitempty"`

	Effect *Effect `json:"effect,omitempty"`
}

func (o Strategy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Strategy struct{}"
	}

	return strings.Join([]string{"Strategy", string(data)}, " ")
}
