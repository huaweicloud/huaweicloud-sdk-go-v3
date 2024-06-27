package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Number struct {
}

func (o Number) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Number struct{}"
	}

	return strings.Join([]string{"Number", string(data)}, " ")
}
