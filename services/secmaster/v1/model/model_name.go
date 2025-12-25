package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Name 名称
type Name struct {
}

func (o Name) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Name struct{}"
	}

	return strings.Join([]string{"Name", string(data)}, " ")
}
