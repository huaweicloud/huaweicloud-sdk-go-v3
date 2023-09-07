package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Location struct {
	X string `json:"x"`

	Y string `json:"y"`
}

func (o Location) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Location struct{}"
	}

	return strings.Join([]string{"Location", string(data)}, " ")
}
