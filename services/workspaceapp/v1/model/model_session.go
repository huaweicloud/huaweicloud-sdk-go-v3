package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Session struct {
	Sbc *Sbc `json:"sbc,omitempty"`
}

func (o Session) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Session struct{}"
	}

	return strings.Join([]string{"Session", string(data)}, " ")
}
