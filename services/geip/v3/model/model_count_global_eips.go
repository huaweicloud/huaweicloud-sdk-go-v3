package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CountGlobalEips struct {

	// Global Eip Count
	Count int32 `json:"count"`
}

func (o CountGlobalEips) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountGlobalEips struct{}"
	}

	return strings.Join([]string{"CountGlobalEips", string(data)}, " ")
}
