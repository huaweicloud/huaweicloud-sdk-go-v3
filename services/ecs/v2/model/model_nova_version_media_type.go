package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NovaVersionMediaType
type NovaVersionMediaType struct {

	// 基础类型。
	Base string `json:"base"`

	// 媒体类型。
	Type string `json:"type"`
}

func (o NovaVersionMediaType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaVersionMediaType struct{}"
	}

	return strings.Join([]string{"NovaVersionMediaType", string(data)}, " ")
}
