package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Marker 页面标记。
type Marker struct {
}

func (o Marker) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Marker struct{}"
	}

	return strings.Join([]string{"Marker", string(data)}, " ")
}
