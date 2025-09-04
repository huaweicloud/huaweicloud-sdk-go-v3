package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Promotion struct {
	Type string `json:"type"`

	Id string `json:"id"`
}

func (o Promotion) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Promotion struct{}"
	}

	return strings.Join([]string{"Promotion", string(data)}, " ")
}
