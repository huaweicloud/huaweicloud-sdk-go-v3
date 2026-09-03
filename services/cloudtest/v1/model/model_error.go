package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Error struct {
	Code *string `json:"code,omitempty"`

	Details *[]interface{} `json:"details,omitempty"`

	Reason *string `json:"reason,omitempty"`

	Url *string `json:"url,omitempty"`
}

func (o Error) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Error struct{}"
	}

	return strings.Join([]string{"Error", string(data)}, " ")
}
