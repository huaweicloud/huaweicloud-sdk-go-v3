package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Destination 目的地址。
type Destination struct {

	// 目的地址。
	Destination string `json:"destination"`
}

func (o Destination) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Destination struct{}"
	}

	return strings.Join([]string{"Destination", string(data)}, " ")
}
