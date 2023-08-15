package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SchedulerAffinity struct {
	Az *[]string `json:"az,omitempty"`

	Node *[]string `json:"node,omitempty"`

	Application *[]string `json:"application,omitempty"`
}

func (o SchedulerAffinity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SchedulerAffinity struct{}"
	}

	return strings.Join([]string{"SchedulerAffinity", string(data)}, " ")
}
