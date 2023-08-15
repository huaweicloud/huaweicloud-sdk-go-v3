package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResQuest struct {
	Limits *Res `json:"limits,omitempty"`

	Requests *Res `json:"requests,omitempty"`
}

func (o ResQuest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResQuest struct{}"
	}

	return strings.Join([]string{"ResQuest", string(data)}, " ")
}
